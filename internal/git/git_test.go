package git

import (
	"os"
	"os/exec"
	"testing"
)

func setupTestRepo(t *testing.T) string {
	tempDir, err := os.MkdirTemp("", "git-test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}

	err = os.Chdir(tempDir)
	if err != nil {
		t.Fatalf("failed to chdir to temp dir: %v", err)
	}

	cmd := exec.Command("git", "init")
	if err := cmd.Run(); err != nil {
		t.Fatalf("failed to init git repo: %v", err)
	}

	return tempDir
}

func TestIsRepo(t *testing.T) {
	oldWd, _ := os.Getwd()
	defer os.Chdir(oldWd)

	tempDir := setupTestRepo(t)
	defer os.RemoveAll(tempDir)

	if !IsRepo() {
		t.Errorf("IsRepo() returned false for a valid repo")
	}

	err := os.Chdir("..")
	if err != nil {
		t.Fatalf("failed to move up: %v", err)
	}
	// In the parent of tempDir, it might still be a repo if the parent is a repo.
	// But in a fresh temp dir, it should be fine.
}

func TestGetOriginURL(t *testing.T) {
	oldWd, _ := os.Getwd()
	defer os.Chdir(oldWd)

	tempDir := setupTestRepo(t)
	defer os.RemoveAll(tempDir)

	url := "https://github.com/test/repo.git"
	cmd := exec.Command("git", "remote", "add", "origin", url)
	if err := cmd.Run(); err != nil {
		t.Fatalf("failed to add remote: %v", err)
	}

	got, err := GetOriginURL()
	if err != nil {
		t.Errorf("GetOriginURL() failed: %v", err)
	}
	if got != url {
		t.Errorf("GetOriginURL() = %v, want %v", got, url)
	}
}
