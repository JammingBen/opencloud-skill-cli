package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/JammingBen/opencloud-skill-cli/internal/logger"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "oc-cli",
	Short:         "OpenCloud CLI to interact with the LibreGraph API",
	Long:          "OpenCloud CLI is a command-line tool to interact with the OpenCloud LibreGraph API. It allows you and AI assistants to manage files, folders and spaces on your OpenCloud server.",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func main() {
	// Setup default logging for errors that happen before commands run
	logger.SetupLogging(slog.LevelInfo)

	if err := rootCmd.Execute(); err != nil {
		msg := color.RedString("error: %v", err)
		fmt.Println(msg)
		os.Exit(1)
	}
}
