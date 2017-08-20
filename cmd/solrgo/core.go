package main

import (
	"context"
	"os"

	"github.com/at15/go-solr/pkg"
	"github.com/at15/go-solr/pkg/common"
	"github.com/spf13/cobra"
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
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("must provide core name")
			return
		}
		name := args[0]
		// TODO: the create solr client part could be shared by most commands
		c := pkg.Config{}
		if addr := os.Getenv("GO_SOLR_ADDR"); addr != "" {
			log.Infof("solr addr %s set via env", addr)
			c.Addr = addr
		}
		solr, err := pkg.New(c)
		if err != nil {
			log.Fatal(err)
			return
		}
		if err := solr.Admin.CreateCoreIfNotExists(context.Background(), common.NewCore(name)); err != nil {
			log.Fatalf("Create core %s failed %v", name, err)
			return
		} else {
			log.Infof("Created core %s (or it already exists)", name)
		}
	},
}

var CoreDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete Solr core",
	Long:  "Delete Solr core using default managed schema",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("let's delete!")
	},
}

func init() {
	CoreCmd.AddCommand(CoreCreateCmd)
	CoreCmd.AddCommand(CoreDeleteCmd)
}
