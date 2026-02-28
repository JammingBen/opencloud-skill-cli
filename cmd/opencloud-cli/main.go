package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "oc-cli",
	Short: "OpenCloud CLI to interact with the LibreGraph API",
	Long:  "OpenCloud CLI is a command-line tool to interact with the OpenCloud LibreGraph API. It allows you and AI assistants to manage files, folders and spaces on your OpenCloud server.",
}

func main() {
	rootCmd.Execute()
}
