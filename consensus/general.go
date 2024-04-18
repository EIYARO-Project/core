package consensus

import (
	"encoding/binary"
	"math"
	"strings"

	"core/protocol/bc"
)

// consensus variables
const (
	// MaxBlockGas Max gas that one block contains
	MaxBlockGas      = uint64(10000000)
	VMGasRate        = int64(200)
	StorageGasRate   = int64(1)
	MaxGasAmount     = int64(200000)
	DefaultGasCredit = int64(30000)

	// CoinbasePendingBlockNumber config parameter for coinbase reward
	CoinbasePendingBlockNumber = uint64(50)
	subsidyReductionInterval   = uint64(175200)
	baseSubsidy                = uint64(100000000000)
	InitialBlockSubsidy        = uint64(21000000000000000)
	subsidyReductionRate       = 0.1

	// BlocksPerRetarget config for pow mining
	BlocksPerRetarget     = uint64(1000)
	TargetSecondsPerBlock = uint64(180)
	SeedPerRetarget       = uint64(256)

	// MaxTimeOffsetSeconds is the maximum number of seconds a block time is allowed to be ahead of the current time
	MaxTimeOffsetSeconds = uint64(10 * 60)
	MedianTimeBlocks     = 8

	PayToWitnessPubKeyHashDataSize = 20
	PayToWitnessScriptHashDataSize = 32
	CoinbaseArbitrarySizeLimit     = 128

	EYAlias = "EY"
)

// EYAssetID is EY's asset id, the soul asset of Eiyaro
var EYAssetID = &bc.AssetID{
	V0: binary.BigEndian.Uint64([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}),
	V1: binary.BigEndian.Uint64([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}),
	V2: binary.BigEndian.Uint64([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}),
	V3: binary.BigEndian.Uint64([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}),
}

// InitialSeed is SHA3-256 of Byte[0^32]
var InitialSeed = &bc.Hash{
	V0: uint64(11412844483649490393),
	V1: uint64(4614157290180302959),
	V2: uint64(1780246333311066183),
	V3: uint64(9357197556716379726),
}

// EYDefinitionMap is the ....
var EYDefinitionMap = map[string]interface{}{
	"name":        EYAlias,
	"symbol":      EYAlias,
	"decimals":    8,
	"description": `Eiyaro Official Issue`,
}

// BlockSubsidy calculates the coinbase reward for a given block height, considering a 10% reduction every subsidyReductionInterval blocks.
func BlockSubsidy(height uint64) uint64 {
	if height == 0 {
		return InitialBlockSubsidy
	}

	// Determine the number of full subsidy reduction intervals that have passed
	intervalsPassed := height / subsidyReductionInterval

	// Calculate the subsidy reduction factor based on the number of intervals passed
	reductionFactor := math.Pow(1-subsidyReductionRate, float64(intervalsPassed))

	// Apply the reduction factor to the base subsidy
	subsidy := uint64(math.Round(float64(baseSubsidy) * reductionFactor))

	return subsidy
}

// IsBech32SegwitPrefix returns whether the prefix is a known prefix for segwit
// addresses on any default or registered network.  This is used when decoding
// an address string into a specific address type.
func IsBech32SegwitPrefix(prefix string, params *Params) bool {
	prefix = strings.ToLower(prefix)
	return prefix == params.Bech32HRPSegwit+"1"
}

// Checkpoint identifies a known good point in the block chain.  Using
// checkpoints allows a few optimizations for old blocks during initial download
// and also prevents forks from old blocks.
type Checkpoint struct {
	Height uint64
	Hash   bc.Hash
}

// Params store the config for different network
type Params struct {
	// Name defines a human-readable identifier for the network.
	Name            string
	Bech32HRPSegwit string
	// DefaultPort defines the default peer-to-peer port for the network.
	DefaultPort string

	// DNSSeeds defines a list of DNS seeds for the network that are used
	// as one method to discover peers.
	DNSSeeds    []string
	Checkpoints []Checkpoint
}

// ActiveNetParams is ...
var ActiveNetParams = MainNetParams

// NetParams is the correspondence between chain_id and Params
var NetParams = map[string]Params{
	"mainnet": MainNetParams,
	"wisdom":  TestNetParams,
	"solonet": SoloNetParams,
}

// MainNetParams is the config for production
var MainNetParams = Params{
	Name:            "main",
	Bech32HRPSegwit: "ey",
	DefaultPort:     "46657",
	DNSSeeds:        []string{"mainnetseed.eiyaro.org"},
	Checkpoints:     []Checkpoint{},
}

// TestNetParams is the config for test-net
var TestNetParams = Params{
	Name:            "test",
	Bech32HRPSegwit: "ty",
	DefaultPort:     "46656",
	DNSSeeds:        []string{"testnetseed.eiyaro.org"},
	Checkpoints:     []Checkpoint{},
}

// SoloNetParams is the config for test-net
var SoloNetParams = Params{
	Name:            "solo",
	Bech32HRPSegwit: "sy",
	Checkpoints:     []Checkpoint{},
}
