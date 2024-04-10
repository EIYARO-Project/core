package api

import (
	"context"

	"core/account"
	"core/asset"
	"core/blockchain/pseudohsm"
	"core/blockchain/rpc"
	"core/blockchain/signers"
	"core/blockchain/txbuilder"
	"core/errors"
	"core/net/http/httperror"
	"core/net/http/httpjson"
	"core/protocol/validation"
	"core/protocol/vm"
)

var (
	// ErrDefault is default Eiyaro API Error
	ErrDefault = errors.New("Eiyaro API Error")
)

func isTemporary(info httperror.Info, err error) bool {
	switch info.ChainCode {
	case "EY000": // internal server error
		return true
	case "EY001": // request timed out
		return true
	case "EY761": // outputs currently reserved
		return true
	case "EY706": // 1 or more action errors
		errs := errors.Data(err)["actions"].([]httperror.Response)
		temp := true
		for _, actionErr := range errs {
			temp = temp && isTemporary(actionErr.Info, nil)
		}
		return temp
	default:
		return false
	}
}

var respErrFormatter = map[error]httperror.Info{
	ErrDefault: {500, "EY000", "Eiyaro API Error"},

	// Signers error namespace (2xx)
	signers.ErrBadQuorum: {400, "EY200", "Quorum must be greater than or equal to 1, and must be less than or equal to the length of xpubs"},
	signers.ErrBadXPub:   {400, "EY201", "Invalid xpub format"},
	signers.ErrNoXPubs:   {400, "EY202", "At least one xpub is required"},
	signers.ErrDupeXPub:  {400, "EY203", "Root XPubs cannot contain the same key more than once"},

	// Contract error namespace (3xx)
	ErrCompileContract: {400, "EY300", "Compile contract failed"},
	ErrInstContract:    {400, "EY301", "Instantiate contract failed"},

	// Transaction error namespace (7xx)
	// Build transaction error namespace (70x ~ 72x)
	account.ErrInsufficient:         {400, "EY700", "Funds of account are insufficient"},
	account.ErrImmature:             {400, "EY701", "Available funds of account are immature"},
	account.ErrReserved:             {400, "EY702", "Available UTXOs of account have been reserved"},
	account.ErrMatchUTXO:            {400, "EY703", "UTXO with given hash not found"},
	ErrBadActionType:                {400, "EY704", "Invalid action type"},
	ErrBadAction:                    {400, "EY705", "Invalid action object"},
	ErrBadActionConstruction:        {400, "EY706", "Invalid action construction"},
	txbuilder.ErrMissingFields:      {400, "EY707", "One or more fields are missing"},
	txbuilder.ErrBadAmount:          {400, "EY708", "Invalid asset amount"},
	account.ErrFindAccount:          {400, "EY709", "Account not found"},
	asset.ErrFindAsset:              {400, "EY710", "Asset not found"},
	txbuilder.ErrBadContractArgType: {400, "EY711", "Invalid contract argument type"},
	txbuilder.ErrOrphanTx:           {400, "EY712", "Transaction input UTXO not found"},
	txbuilder.ErrExtTxFee:           {400, "EY713", "Transaction fee exceeded max limit"},
	txbuilder.ErrNoGasInput:         {400, "EY714", "Transaction has no gas input"},

	// Submit transaction error namespace (73x ~ 79x)
	// Validation error (73x ~ 75x)
	validation.ErrTxVersion:                 {400, "EY730", "Invalid transaction version"},
	validation.ErrWrongTransactionSize:      {400, "EY731", "Invalid transaction size"},
	validation.ErrBadTimeRange:              {400, "EY732", "Invalid transaction time range"},
	validation.ErrNotStandardTx:             {400, "EY733", "Not standard transaction"},
	validation.ErrWrongCoinbaseTransaction:  {400, "EY734", "Invalid coinbase transaction"},
	validation.ErrWrongCoinbaseAsset:        {400, "EY735", "Invalid coinbase assetID"},
	validation.ErrCoinbaseArbitraryOversize: {400, "EY736", "Invalid coinbase arbitrary size"},
	validation.ErrEmptyResults:              {400, "EY737", "No results in the transaction"},
	validation.ErrMismatchedAssetID:         {400, "EY738", "Mismatched assetID"},
	validation.ErrMismatchedPosition:        {400, "EY739", "Mismatched value source/dest position"},
	validation.ErrMismatchedReference:       {400, "EY740", "Mismatched reference"},
	validation.ErrMismatchedValue:           {400, "EY741", "Mismatched value"},
	validation.ErrMissingField:              {400, "EY742", "Missing required field"},
	validation.ErrNoSource:                  {400, "EY743", "No source for value"},
	validation.ErrOverflow:                  {400, "EY744", "Arithmetic overflow/underflow"},
	validation.ErrPosition:                  {400, "EY745", "Invalid source or destination position"},
	validation.ErrUnbalanced:                {400, "EY746", "Unbalanced asset amount between input and output"},
	validation.ErrOverGasCredit:             {400, "EY747", "Gas credit has been spent"},
	validation.ErrGasCalculate:              {400, "EY748", "Gas usage calculate got a math error"},

	// VM error (76x ~ 78x)
	vm.ErrAltStackUnderflow:  {400, "EY760", "Alt stack underflow"},
	vm.ErrBadValue:           {400, "EY761", "Bad value"},
	vm.ErrContext:            {400, "EY762", "Wrong context"},
	vm.ErrDataStackUnderflow: {400, "EY763", "Data stack underflow"},
	vm.ErrDisallowedOpcode:   {400, "EY764", "Disallowed opcode"},
	vm.ErrDivZero:            {400, "EY765", "Division by zero"},
	vm.ErrFalseVMResult:      {400, "EY766", "False result for executing VM"},
	vm.ErrLongProgram:        {400, "EY767", "Program size exceeds max int32"},
	vm.ErrRange:              {400, "EY768", "Arithmetic range error"},
	vm.ErrReturn:             {400, "EY769", "RETURN executed"},
	vm.ErrRunLimitExceeded:   {400, "EY770", "Run limit exceeded because the EY Fee is insufficient"},
	vm.ErrShortProgram:       {400, "EY771", "Unexpected end of program"},
	vm.ErrToken:              {400, "EY772", "Unrecognized token"},
	vm.ErrUnexpected:         {400, "EY773", "Unexpected error"},
	vm.ErrUnsupportedVM:      {400, "EY774", "Unsupported VM because the version of VM is mismatched"},
	vm.ErrVerifyFailed:       {400, "EY775", "VERIFY failed"},

	// Mock HSM error namespace (8xx)
	pseudohsm.ErrDuplicateKeyAlias: {400, "EY800", "Key Alias already exists"},
	pseudohsm.ErrLoadKey:           {400, "EY801", "Key not found or wrong password"},
	pseudohsm.ErrDecrypt:           {400, "EY802", "Could not decrypt key with given passphrase"},
}

// Map error values to standard eiyaro error codes. Missing entries
// will map to internalErrInfo.
//
// TODO(jackson): Share one error table across Chain
// products/services so that errors are consistent.
var errorFormatter = httperror.Formatter{
	Default:     httperror.Info{500, "EY000", "Eiyaro API Error"},
	IsTemporary: isTemporary,
	Errors: map[error]httperror.Info{
		// General error namespace (0xx)
		context.DeadlineExceeded: {408, "EY001", "Request timed out"},
		httpjson.ErrBadRequest:   {400, "EY002", "Invalid request body"},
		rpc.ErrWrongNetwork:      {502, "EY103", "A peer core is operating on a different blockchain network"},

		//accesstoken authz err namespace (86x)
		errNotAuthenticated: {401, "EY860", "Request could not be authenticated"},
	},
}
