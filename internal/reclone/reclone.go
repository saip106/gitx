package reclone

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/user/gitx/internal/fs"
	"github.com/user/gitx/internal/git"
	"github.com/user/gitx/pkg/prompt"
	"github.com/spf13/cobra"
)

var (
	force        bool
	depth        int
	singleBranch bool
	noTags       bool
	verbose      bool
)

// NewRecloneCmd creates the reclone command.
func NewRecloneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reclone [path]",
		Short: "Delete and re-clone a repository",
		Args:  cobra.MaximumNArgs(1),
		RunE:  runReclone,
	}

	cmd.Flags().BoolVarP(&force, "force", "f", false, "Skip all confirmations")
	cmd.Flags().IntVar(&depth, "depth", 0, "Perform a shallow clone")
	cmd.Flags().BoolVar(&singleBranch, "single-branch", false, "Clone only the current branch")
	cmd.Flags().BoolVar(&noTags, "no-tags", false, "Do not fetch tags")
	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Print detailed logs")

	return cmd
}

func runReclone(cmd *cobra.Command, args []string) error {
	targetPath := "."
	if len(args) > 0 {
		targetPath = args[0]
	}

	// Change to target path to detect repo and get info
	if err := os.Chdir(targetPath); err != nil {
		return fmt.Errorf("failed to access path %s: %w", targetPath, err)
	}

	if !git.IsRepo() {
		return fmt.Errorf("directory %s is not a git repository", targetPath)
	}

	absPath, err := filepath.Abs(".")
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}
	folderName := filepath.Base(absPath)
	parentDir := filepath.Dir(absPath)

	if verbose {
		fmt.Printf("Detected repo path: %s\n", absPath)
	}

	hasChanges, err := git.HasUncommittedChanges()
	if err != nil {
		return err
	}

	if verbose {
		fmt.Printf("Pending changes: %v\n", hasChanges)
	}

	if hasChanges && !force {
		fmt.Println("WARNING: You have uncommitted changes.")
		if !prompt.Confirm("Continue and LOSE these changes?", false) {
			fmt.Println("Aborted.")
			return nil
		}
	} else if !force {
		if !prompt.Confirm("This will delete and reclone the repository. Continue?", false) {
			fmt.Println("Aborted.")
			return nil
		}
	}

	originURL, err := git.GetOriginURL()
	if err != nil {
		return err
	}

	if verbose {
		fmt.Printf("Remote URL: %s\n", originURL)
	}

	// Move to parent directory
	if err := os.Chdir(parentDir); err != nil {
		return fmt.Errorf("failed to move to parent directory: %w", err)
	}

	if verbose {
		fmt.Println("Attempting safe delete...")
	}

	err = fs.SafeDelete(absPath)
	if err != nil {
		if verbose {
			fmt.Printf("Safe delete failed: %v\n", err)
		}
		if force || prompt.Confirm("Safe delete failed. Perform hard delete (rm -rf)?", false) {
			if verbose {
				fmt.Println("Performing hard delete...")
			}
			if err := fs.HardDelete(absPath); err != nil {
				return fmt.Errorf("hard delete failed: %w", err)
			}
		} else {
			fmt.Println("Aborted.")
			return nil
		}
	} else if verbose {
		fmt.Println("Safe delete successful.")
	}

	if verbose {
		fmt.Printf("Cloning %s into %s...\n", originURL, folderName)
	}

	if err := git.Clone(originURL, folderName, depth, singleBranch, noTags); err != nil {
		return err
	}

	// Return to the folder
	if err := os.Chdir(folderName); err != nil {
		return fmt.Errorf("failed to return to folder: %w", err)
	}

	fmt.Println("Successfully recloned repository.")
	return nil
}
