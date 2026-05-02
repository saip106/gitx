package git

import (
	"fmt"
	"os/exec"
	"strings"
)

// IsRepo checks if the current directory is a git repository.
func IsRepo() bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

// HasUncommittedChanges checks if there are any uncommitted changes in the repository.
func HasUncommittedChanges() (bool, error) {
	cmd := exec.Command("git", "status", "--porcelain")
	output, err := cmd.Output()
	if err != nil {
		return false, fmt.Errorf("failed to check git status: %w", err)
	}
	return len(strings.TrimSpace(string(output))) > 0, nil
}

// GetOriginURL returns the URL of the 'origin' remote.
func GetOriginURL() (string, error) {
	cmd := exec.Command("git", "remote", "get-url", "origin")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get origin URL: %w", err)
	}
	return strings.TrimSpace(string(output)), nil
}

// Clone clones a repository into a specific directory with optional flags.
func Clone(url, dir string, depth int, singleBranch, noTags bool) error {
	args := []string{"clone"}
	if depth > 0 {
		args = append(args, "--depth", fmt.Sprintf("%d", depth))
	}
	if singleBranch {
		args = append(args, "--single-branch")
	}
	if noTags {
		args = append(args, "--no-tags")
	}
	args = append(args, url, dir)

	cmd := exec.Command("git", args...)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to clone: %s: %w", string(output), err)
	}
	return nil
}
