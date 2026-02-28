package main

import (
	"context"
	"fmt"

	"github.com/JammingBen/opencloud-skill-cli/internal/config"
	"github.com/JammingBen/opencloud-skill-cli/internal/http"
	"github.com/spf13/cobra"
)

var path string
var method string

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Interact with the OpenCloud LibreGraph API",
	Long:  "This command allows you to interact with the OpenCloud LibreGraph API.",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Load config
		cfg, err := config.Load()
		if err != nil {
			return err
		}

		if cfg.ServerURL == "" {
			return fmt.Errorf("Server URL is not configured. Please run 'occ login' first.")
		}

		// Get the token resource for authentication
		ctx := context.Background()
		ts, err := cfg.GetTokenSource(ctx)
		if err != nil {
			return err
		}
		if ts == nil {
			return fmt.Errorf("Access token is not configured. Please run 'occ login' first.")
		}

		// Create client and make request to OpenCloud API
		c := http.NewClient(cfg.ServerURL, cfg.Insecure, ts)
		resp, err := c.MakeRequest(path, method)
		if err != nil {
			return err
		}

		// Print output
		fmt.Println(resp)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
	apiCmd.Flags().StringVarP(&path, "path", "p", "/", "Path of the API endpoint")
	apiCmd.Flags().StringVarP(&method, "method", "m", "GET", "HTTP method to use")
	apiCmd.MarkFlagRequired("path")
}
