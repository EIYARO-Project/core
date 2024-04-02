package commands

import (
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/cobra"

	"eiyaro/util"
)

// eiyarocli usage template
var usageTemplate = `Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

Available Commands:
    {{range .Commands}}{{if (and .IsAvailableCommand (.Name | WalletDisable))}}
    {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}

  available with wallet enable:
    {{range .Commands}}{{if (and .IsAvailableCommand (.Name | WalletEnable))}}
    {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`

// commandError is an error used to signal different error situations in command handling.
type commandError struct {
	s         string
	userError bool
}

func (c commandError) Error() string {
	return c.s
}

func (c commandError) isUserError() bool {
	return c.userError
}

func newUserError(a ...interface{}) commandError {
	return commandError{s: fmt.Sprintln(a...), userError: true}
}

func newSystemError(a ...interface{}) commandError {
	return commandError{s: fmt.Sprintln(a...), userError: false}
}

func newSystemErrorF(format string, a ...interface{}) commandError {
	return commandError{s: fmt.Sprintf(format, a...), userError: false}
}

// Catch some of the obvious user errors from Cobra.
// We don't want to show the usage message for every error.
// The below may be to generic. Time will show.
var userErrorRegexp = regexp.MustCompile("argument|flag|shorthand")

func isUserError(err error) bool {
	if cErr, ok := err.(commandError); ok && cErr.isUserError() {
		return true
	}

	return userErrorRegexp.MatchString(err.Error())
}

// EiyarocliCmd is Eiyarocli's root command.
// Every other command attached to EiyarocliCmd is a child command to it.
var EiyarocliCmd = &cobra.Command{
	Use:   "eiyarocli",
	Short: "Eiyarocli is a commond line client for eiyaro core (a.k.a. eiyarod)",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.SetUsageTemplate(usageTemplate)
			cmd.Usage()
		}
	},
}

// Execute adds all child commands to the root command EiyarocliCmd and sets flags appropriately.
func Execute() {

	AddCommands()
	AddTemplateFunc()

	if _, err := EiyarocliCmd.ExecuteC(); err != nil {
		os.Exit(util.ErrLocalExe)
	}
}

// AddCommands adds child commands to the root command EiyarocliCmd.
func AddCommands() {
	EiyarocliCmd.AddCommand(createAccessTokenCmd)
	EiyarocliCmd.AddCommand(listAccessTokenCmd)
	EiyarocliCmd.AddCommand(deleteAccessTokenCmd)
	EiyarocliCmd.AddCommand(checkAccessTokenCmd)

	EiyarocliCmd.AddCommand(createAccountCmd)
	EiyarocliCmd.AddCommand(deleteAccountCmd)
	EiyarocliCmd.AddCommand(listAccountsCmd)
	EiyarocliCmd.AddCommand(updateAccountAliasCmd)
	EiyarocliCmd.AddCommand(createAccountReceiverCmd)
	EiyarocliCmd.AddCommand(listAddressesCmd)
	EiyarocliCmd.AddCommand(validateAddressCmd)
	EiyarocliCmd.AddCommand(listPubKeysCmd)

	EiyarocliCmd.AddCommand(createAssetCmd)
	EiyarocliCmd.AddCommand(getAssetCmd)
	EiyarocliCmd.AddCommand(listAssetsCmd)
	EiyarocliCmd.AddCommand(updateAssetAliasCmd)

	EiyarocliCmd.AddCommand(getTransactionCmd)
	EiyarocliCmd.AddCommand(listTransactionsCmd)

	EiyarocliCmd.AddCommand(getUnconfirmedTransactionCmd)
	EiyarocliCmd.AddCommand(listUnconfirmedTransactionsCmd)
	EiyarocliCmd.AddCommand(decodeRawTransactionCmd)

	EiyarocliCmd.AddCommand(listUnspentOutputsCmd)
	EiyarocliCmd.AddCommand(listBalancesCmd)

	EiyarocliCmd.AddCommand(rescanWalletCmd)
	EiyarocliCmd.AddCommand(walletInfoCmd)

	EiyarocliCmd.AddCommand(buildTransactionCmd)
	EiyarocliCmd.AddCommand(signTransactionCmd)
	EiyarocliCmd.AddCommand(submitTransactionCmd)
	EiyarocliCmd.AddCommand(estimateTransactionGasCmd)

	EiyarocliCmd.AddCommand(getBlockCountCmd)
	EiyarocliCmd.AddCommand(getBlockHashCmd)
	EiyarocliCmd.AddCommand(getBlockCmd)
	EiyarocliCmd.AddCommand(getBlockHeaderCmd)
	EiyarocliCmd.AddCommand(getDifficultyCmd)
	EiyarocliCmd.AddCommand(getHashRateCmd)

	EiyarocliCmd.AddCommand(createKeyCmd)
	EiyarocliCmd.AddCommand(deleteKeyCmd)
	EiyarocliCmd.AddCommand(listKeysCmd)
	EiyarocliCmd.AddCommand(updateKeyAliasCmd)
	EiyarocliCmd.AddCommand(resetKeyPwdCmd)
	EiyarocliCmd.AddCommand(checkKeyPwdCmd)

	EiyarocliCmd.AddCommand(signMsgCmd)
	EiyarocliCmd.AddCommand(verifyMsgCmd)
	EiyarocliCmd.AddCommand(decodeProgCmd)

	EiyarocliCmd.AddCommand(createTransactionFeedCmd)
	EiyarocliCmd.AddCommand(listTransactionFeedsCmd)
	EiyarocliCmd.AddCommand(deleteTransactionFeedCmd)
	EiyarocliCmd.AddCommand(getTransactionFeedCmd)
	EiyarocliCmd.AddCommand(updateTransactionFeedCmd)

	EiyarocliCmd.AddCommand(isMiningCmd)
	EiyarocliCmd.AddCommand(setMiningCmd)

	EiyarocliCmd.AddCommand(netInfoCmd)
	EiyarocliCmd.AddCommand(gasRateCmd)

	EiyarocliCmd.AddCommand(versionCmd)
}

// AddTemplateFunc adds usage template to the root command EiyarocliCmd.
func AddTemplateFunc() {
	walletEnableCmd := []string{
		createAccountCmd.Name(),
		listAccountsCmd.Name(),
		deleteAccountCmd.Name(),
		updateAccountAliasCmd.Name(),
		createAccountReceiverCmd.Name(),
		listAddressesCmd.Name(),
		validateAddressCmd.Name(),
		listPubKeysCmd.Name(),

		createAssetCmd.Name(),
		getAssetCmd.Name(),
		listAssetsCmd.Name(),
		updateAssetAliasCmd.Name(),

		createKeyCmd.Name(),
		deleteKeyCmd.Name(),
		listKeysCmd.Name(),
		resetKeyPwdCmd.Name(),
		checkKeyPwdCmd.Name(),
		signMsgCmd.Name(),

		buildTransactionCmd.Name(),
		signTransactionCmd.Name(),

		getTransactionCmd.Name(),
		listTransactionsCmd.Name(),
		listUnspentOutputsCmd.Name(),
		listBalancesCmd.Name(),

		rescanWalletCmd.Name(),
		walletInfoCmd.Name(),
	}

	cobra.AddTemplateFunc("WalletEnable", func(cmdName string) bool {
		for _, name := range walletEnableCmd {
			if name == cmdName {
				return true
			}
		}
		return false
	})

	cobra.AddTemplateFunc("WalletDisable", func(cmdName string) bool {
		for _, name := range walletEnableCmd {
			if name == cmdName {
				return false
			}
		}
		return true
	})
}
