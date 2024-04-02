package commands

import (
	"runtime"

	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"

	"eiyaro/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Eiyarocli",
	Run: func(cmd *cobra.Command, args []string) {
		jww.FEEDBACK.Printf("Eiyarocli v%s %s/%s\n", version.Version, runtime.GOOS, runtime.GOARCH)
	},
}
