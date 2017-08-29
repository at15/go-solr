package main

import (
	"context"

	"github.com/at15/go-solr/solr"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	overwrite = false // TODO: by default we use create if not exist, but we can delete the old one and create a fresh new one
	configSet = ""
)

var CoreCmd = &cobra.Command{
	Use:   "core",
	Short: "Manage Solr core",
	Long:  "Create, delete, check status of Solr core",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

var CoreCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Solr core",
	Long:  "Create Solr core using default managed schema",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("must provide core name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		core := solr.NewCore(name)
		core.ConfigSet = configSet
		exists, err := solrClient.CreateCoreIfNotExists(context.Background(), core)
		if err != nil {
			log.Fatalf("Create core %s failed %v", name, err)
			return
		}
		if exists {
			log.Infof("Core %s already exists", name)
		} else {
			log.Infof("Created core %s", name)
		}
	},
}

var CoreDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete Solr core",
	Long:  "Delete Solr core using default managed schema",
	Args: func(cmd *cobra.Command, args []string) error {
		// TODO: might put it in parent command's persistent pre run, though we should be checking args[1] instead of args[0] I suppose
		if len(args) < 1 {
			return errors.New("must provide core name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		if err := solrClient.DeleteCore(context.Background(), name); err != nil {
			log.Fatalf("Delete core %s failed %v", name, err)
		} else {
			log.Infof("Core %s deleted", name)
		}
	},
}

func init() {
	CoreCreateCmd.Flags().StringVar(&configSet, "configSet", solr.DefaultConfigSet,
		"specify configSet for the core, it must already exists, you should NOT use the default value if you have more than one cores with different schema")

	CoreCmd.AddCommand(CoreCreateCmd)
	CoreCmd.AddCommand(CoreDeleteCmd)
}
