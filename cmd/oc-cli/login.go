package main

import (
	"github.com/JammingBen/opencloud-skill-cli/internal/oidc"
	"github.com/spf13/cobra"
)

var loginServerUrl string
var loginInsecure bool
var loginClientID string

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to the OpenCloud server via OIDC",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := oidc.NewOIDCClient(loginServerUrl, loginInsecure, loginClientID)
		if err := c.Login(); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringVarP(&loginServerUrl, "server-url", "s", "https://localhost:9200", "URL of the OpenCloud server")
	loginCmd.Flags().BoolVarP(&loginInsecure, "insecure", "k", false, "Allow insecure server connections when using SSL")
	loginCmd.Flags().StringVarP(&loginClientID, "client-id", "i", "", "OAuth2 Client ID (default: discovered from config.json or OpenCloudDesktop)")
	loginCmd.MarkFlagRequired("server-url")
}
