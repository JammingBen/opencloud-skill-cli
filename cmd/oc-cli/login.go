package main

import (
	"github.com/JammingBen/opencloud-skill-cli/internal/oidc"
	"github.com/spf13/cobra"
)

var (
	serverUrl string
	insecure  bool
	clientID  string
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to the OpenCloud server via OIDC",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := oidc.NewOIDCClient(serverUrl, insecure, clientID)
		if err := c.Login(); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringVarP(&serverUrl, "server-url", "s", "https://localhost:9200", "URL of the OpenCloud server")
	loginCmd.Flags().BoolVarP(&insecure, "insecure", "k", false, "Allow insecure server connections when using SSL")
	loginCmd.Flags().StringVarP(&clientID, "client-id", "i", "", "OAuth2 Client ID (default: discovered from config.json or OpenCloudDesktop)")
	loginCmd.MarkFlagRequired("server-url")
}
