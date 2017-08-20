package main

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	verbose = false
	yes     = false
)

var RootCmd = &cobra.Command{
	Use:   "gosolr",
	Short: "Solr client in go",
	Long:  "Solr client and util in go https://github.com/at15/go-solr",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func main() {
	if RootCmd.Execute() != nil {
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.AddCommand(VersionCmd)
	RootCmd.AddCommand(CoreCmd)

	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable debug logging and print full information")
	RootCmd.PersistentFlags().BoolVarP(&yes, "yes", "y", false, "yes to all prompt")
}

func initConfig() {
	if verbose {
		Logger.SetLevel("debug")
		// NOTE: this is disabled because version command use --verbose as well and we just want to see the version output
		//log.Debug("enabled debug level logging")
	}
}
