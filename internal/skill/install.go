package skill

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/JammingBen/opencloud-skill-cli/internal/version"
)

const CLAUDE_CODE_DIR = ".claude/skills"
const GITHUB_COPILOT_SKILL_DIR = ".copilot/skills"
const CODEX_SKILL_DIR = ".agents/skills"
const OPEN_CODE_SKILL_DIR = ".config/opencode/skills"
const GEMINI_SKILL_DIR = ".gemini/skills"

// InstallSkill downloads the skill zip from GitHub, extracts it, and saves it to the appropriate directory based on the agent
func InstallSkill(agent string) error {
	c := http.Client{}

	req, _ := http.NewRequest("GET", getZipUrl(), nil)
	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download skill: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		return fmt.Errorf("failed to create zip reader: %w", err)
	}

	for _, f := range zipReader.File {
		path, err := getSkillDirPath(agent, f)
		if err != nil {
			return err
		}

		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(path, os.ModePerm); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", path, err)
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
			return fmt.Errorf("failed to create parent directory for %s: %w", path, err)
		}

		dstFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return fmt.Errorf("failed to open destination file %s: %w", path, err)
		}

		srcFile, err := f.Open()
		if err != nil {
			dstFile.Close()
			return fmt.Errorf("failed to open file in zip %s: %w", f.Name, err)
		}

		_, err = io.Copy(dstFile, srcFile)
		srcFile.Close()
		dstFile.Close()
		if err != nil {
			return fmt.Errorf("failed to copy content to %s: %w", path, err)
		}
	}

	return nil
}

// getZipUrl constructs the URL to download the skill zip based on the current version
func getZipUrl() string {
	v := version.GetVersionWithoutSuffix()
	return fmt.Sprintf("https://github.com/JammingBen/opencloud-skill-cli/releases/download/v%s/skill.zip", v)
}

// getSkillDirPath returns the appropriate skill directory path based on the agent and the file in the zip
func getSkillDirPath(agent string, f *zip.File) (string, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}

	var skillDir string
	switch agent {
	case "claude-code":
		skillDir = filepath.Join(userHomeDir, CLAUDE_CODE_DIR, f.Name)
	case "github-copilot":
		skillDir = filepath.Join(userHomeDir, GITHUB_COPILOT_SKILL_DIR, f.Name)
	case "codex":
		skillDir = filepath.Join(userHomeDir, CODEX_SKILL_DIR, f.Name)
	case "open-code":
		skillDir = filepath.Join(userHomeDir, OPEN_CODE_SKILL_DIR, f.Name)
	case "gemini-cli":
		skillDir = filepath.Join(userHomeDir, GEMINI_SKILL_DIR, f.Name)
	default:
		return "", fmt.Errorf("unsupported agent: %s", agent)
	}

	return skillDir, nil
}
