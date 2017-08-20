package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// injected at build time using ldflag, see Makefile for detail
var (
	Version     string
	BuildUser   string
	BuildCommit string
	BuildTime   string
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Long:  "Show version, git commit and build time (if set)",
	Run: func(cmd *cobra.Command, args []string) {
		if !verbose {
			fmt.Println(Version)
			return
		}
		fmt.Printf("version: %s\n", Version)
		fmt.Printf("build commit: %s\n", BuildCommit)
		fmt.Printf("build time: %s\n", BuildTime)
		fmt.Printf("build user: %s\n", BuildUser)
	},
}
