package main

import (
	"fmt"

	"github.com/JammingBen/opencloud-skill-cli/internal/oidc"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout from the OpenCloud server",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := oidc.LoadConfig()
		if err != nil {
			return err
		}

		if err := cfg.Clear(); err != nil {
			return fmt.Errorf("failed to clear session: %w", err)
		}

		fmt.Println("Successfully logged out and cleared session.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
