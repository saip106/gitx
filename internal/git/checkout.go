package git

import (
	"fmt"
	"os/exec"
)

// CheckoutAndPull switches to the given branch and runs 'git pull --all'.
// Returns combined output and error (if any).
func CheckoutAndPull(branch string) (string, error) {
	// Switch branch
	checkout := exec.Command("git", "checkout", branch)
	out, err := checkout.CombinedOutput()
	if err != nil {
		return string(out), fmt.Errorf("failed to checkout branch '%s': %w", branch, err)
	}
	// Pull all
	pull := exec.Command("git", "pull", "--all")
	pullOut, pullErr := pull.CombinedOutput()
	if pullErr != nil {
		return string(out) + string(pullOut), fmt.Errorf("failed to pull all after checkout: %w", pullErr)
	}
	return string(out) + string(pullOut), nil
}
