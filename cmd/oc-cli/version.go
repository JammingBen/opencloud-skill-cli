package main

import (
	"fmt"

	"github.com/JammingBen/opencloud-skill-cli/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of OpenCloud CLI",
	Long:  "All software has versions. This is OpenCloud CLI's version.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("v%s\n", version.GetVersion())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
