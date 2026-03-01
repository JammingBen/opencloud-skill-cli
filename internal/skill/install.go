package skill

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

const CLAUDE_CODE_DIR = ".claude/skills"
const GITHUB_COPILOT_SKILL_DIR = ".copilot/skills"
const CODEX_SKILL_DIR = ".agents/skills"
const OPEN_CODE_SKILL_DIR = ".config/opencode/skills"
const GEMINI_SKILL_DIR = ".gemini/skills"

// InstallSkill installs the skill from the embedded filesystem for the appropriate directory based on the agent
func InstallSkill(agent string, embedFS fs.FS) error {
	destDir, err := getAgentSkillDir(agent)
	if err != nil {
		return err
	}

	// The source path "skills" matches //go:embed directive in embed.go
	return CopyFS(embedFS, "skills", destDir)
}

// CopyFS copies an embedded filesystem to a local directory
func CopyFS(embedFS fs.FS, srcDir, destDir string) error {
	return fs.WalkDir(embedFS, srcDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Determine the relative path from the source directory
		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}

		targetPath := filepath.Join(destDir, relPath)

		if d.IsDir() {
			return os.MkdirAll(targetPath, 0755)
		}

		// Read the file from the embedded FS
		data, err := fs.ReadFile(embedFS, path)
		if err != nil {
			return err
		}

		// Ensure parent directory exists (already covered by d.IsDir() check for most cases, but safe for root files)
		if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
			return err
		}

		// Write the file to the local FS
		return os.WriteFile(targetPath, data, 0644)
	})
}

// getAgentSkillDir returns the appropriate skill directory base path based on the agent
func getAgentSkillDir(agent string) (string, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}

	var skillDir string
	switch agent {
	case "claude-code":
		skillDir = filepath.Join(userHomeDir, CLAUDE_CODE_DIR)
	case "github-copilot":
		skillDir = filepath.Join(userHomeDir, GITHUB_COPILOT_SKILL_DIR)
	case "codex":
		skillDir = filepath.Join(userHomeDir, CODEX_SKILL_DIR)
	case "open-code":
		skillDir = filepath.Join(userHomeDir, OPEN_CODE_SKILL_DIR)
	case "gemini-cli":
		skillDir = filepath.Join(userHomeDir, GEMINI_SKILL_DIR)
	default:
		return "", fmt.Errorf("unsupported agent: %s", agent)
	}

	return skillDir, nil
}
