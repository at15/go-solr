package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "gosolr",
	Short: "Solr client in go",
	Long:  "Solr client and util in go https://github.com/at15/go-solr",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hi")
	},
}

func main() {
	if RootCmd.Execute() != nil {
		os.Exit(-1)
	}
}
