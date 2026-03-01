package main

import (
	"fmt"

	opencloudskillcli "github.com/JammingBen/opencloud-skill-cli"
	"github.com/JammingBen/opencloud-skill-cli/internal/skill"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var agent string

var installSkillCmd = &cobra.Command{
	Use:   "install-skill",
	Short: "Install a new OpenCloud CLI skill",
	Long:  "This command installs the OpenCloud CLI skill for AI assistants and agents.",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := skill.InstallSkill(agent, opencloudskillcli.SkillFolder)
		if err != nil {
			return fmt.Errorf("failed to install skill: %w", err)
		}

		msg := color.GreenString("✓ Successfully installed skill for %s", agent)
		fmt.Println(msg)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(installSkillCmd)
	installSkillCmd.Flags().StringVarP(&agent, "agent", "a", "", "Agent to install the skill for (claude-code, github-copilot, codex, open-code, gemini-cli)")
	installSkillCmd.MarkFlagRequired("agent")
}
