package fs

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// SafeDelete attempts to move the directory to the trash/recycle bin.
func SafeDelete(path string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		// Use PowerShell to move to Recycle Bin
		psCommand := fmt.Sprintf("Add-Type -AssemblyName Microsoft.VisualBasic; [Microsoft.VisualBasic.FileIO.FileSystem]::DeleteDirectory('%s', 'OnlyErrorDialogs', 'SendToRecycleBin')", path)
		cmd = exec.Command("powershell", "-NoProfile", "-Command", psCommand)
	case "darwin":
		// Use osascript to move to Trash
		osaCommand := fmt.Sprintf("tell application \"Finder\" to delete POSIX file \"%s\"", path)
		cmd = exec.Command("osascript", "-e", osaCommand)
	case "linux":
		// Fallback for Linux - move to ~/.local/share/Trash/files if it exists
		trashPath := os.Getenv("HOME") + "/.local/share/Trash/files"
		if _, err := os.Stat(trashPath); err == nil {
			// Simple move for now, though a real trash implementation is more complex (metadata)
			// But for this tool, a move to a known trash dir is better than nothing.
			// However, most linux users expect 'trash-cli' or similar.
			// Let's just try to move it to a subfolder there.
			dest := trashPath + "/" + path
			return os.Rename(path, dest)
		}
		return fmt.Errorf("safe delete not supported on this linux configuration")
	default:
		return fmt.Errorf("safe delete not supported on %s", runtime.GOOS)
	}

	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("safe delete failed: %s: %w", string(output), err)
	}
	return nil
}

// HardDelete performs a hard delete (rm -rf).
func HardDelete(path string) error {
	return os.RemoveAll(path)
}
