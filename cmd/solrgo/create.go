package main

import (
	"github.com/spf13/cobra"
	"github.com/at15/go-solr/solr"
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create core etc.",
	Long:  "Create core etc.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

// CreateCoreCmd is an alias for CoreCreateCmd, so users can use solrgo create core demo instead solrgo core create demo
// TODO: though I am quite wondering how golang determine the initialization order of package level variables
var CreateCoreCmd = &cobra.Command{
	Use:   "core",
	Short: CoreCreateCmd.Short,
	Long:  CoreCreateCmd.Long,
	Args:  CoreCreateCmd.Args,
	Run:   CoreCreateCmd.Run,
}

func init() {
	//coreCreateCmd := new(cobra.Command)
	//*coreCreateCmd = *CoreCreateCmd
	// FIXME: it seems change the use does not works, we still see create in available commands
	//coreCreateCmd.Use = "core"
	// FIXME: we are copying from CoreCreateCmd due to we can't just change the use ...
	CreateCoreCmd.Flags().StringVar(&configSet, "configSet", solr.DefaultConfigSet,
		"specify configSet for the core, it must already exists, you should NOT use the default value if you have more than one cores with different schema")

	CreateCmd.AddCommand(CreateCoreCmd)
}
