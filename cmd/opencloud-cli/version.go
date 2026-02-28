package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Version is the version of the OpenCloud CLI, which can be set at build time via ldflags.
	Version = "0.0.1"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of OpenCloud CLI",
	Long:  "All software has versions. This is OpenCloud CLI's version.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("v%s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
