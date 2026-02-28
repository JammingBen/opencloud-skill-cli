package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var installSkillCmd = &cobra.Command{
	Use:   "install-skill",
	Short: "Install a new OpenCloud CLI skill",
	Long:  "This command installs the OpenCloud CLI skill for AI assistants and agents.",
	Run: func(cmd *cobra.Command, args []string) {
		// FIXME: This should be implemented to install a skill
		fmt.Println("Installing a new skill...")
	},
}

func init() {
	rootCmd.AddCommand(installSkillCmd)
}
