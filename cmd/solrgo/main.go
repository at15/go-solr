package main

import (
	"os"

	"github.com/at15/go-solr/solr"
	"github.com/spf13/cobra"
)

var (
	verbose    = false
	yes        = false
	solrClient *solr.Client
)

var RootCmd = &cobra.Command{
	Use:   "gosolr",
	Short: "Solr client in go",
	Long:  "Solr client and util in go https://github.com/at15/go-solr",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// skip initialize for version command
		if cmd.Use == "version" {
			return
		}
		initConfig()
	},
}

func main() {
	if RootCmd.Execute() != nil {
		os.Exit(-1)
	}
}

func init() {
	RootCmd.AddCommand(VersionCmd)
	RootCmd.AddCommand(CreateCmd)
	RootCmd.AddCommand(CoreCmd)

	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable debug logging and print full information")
	RootCmd.PersistentFlags().BoolVarP(&yes, "yes", "y", false, "yes to all prompt")
}

func initConfig() {
	if verbose {
		Logger.SetLevel("debug")
		log.Debug("enabled debug level logging")
	}
	var err error
	c := solr.Config{
		Addr: os.Getenv(solr.AddrEnvName),
	}
	if solrClient, err = solr.NewClient(c); err != nil {
		log.Fatalf("can't initial solr client %v", err)
	}
}
