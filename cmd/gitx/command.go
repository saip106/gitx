package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/user/gitx/internal/git"
)

// NewMasterCmd creates the 'master' command.
func NewMasterCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "master",
		Short: "Switch to master and pull all remotes",
		Long: `Switches to the 'master' branch and runs 'git pull --all'.

	Shows all git output and a summary. Fails if the branch does not exist.
	Useful for quickly updating your local master branch from all remotes.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			output, err := git.CheckoutAndPull("master")
			fmt.Print(output)
			if err != nil {
				return fmt.Errorf("[gitx master] failed: %w", err)
			}
			fmt.Println("\n[gitx master] Switched to 'master' and pulled all remotes.")
			return nil
		},
	}
}

// NewMainCmd creates the 'main' command.
func NewMainCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "main",
		Short: "Switch to main and pull all remotes",
		Long: `Switches to the 'main' branch and runs 'git pull --all'.

	Shows all git output and a summary. Fails if the branch does not exist.
	Useful for quickly updating your local main branch from all remotes.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			output, err := git.CheckoutAndPull("main")
			fmt.Print(output)
			if err != nil {
				return fmt.Errorf("[gitx main] failed: %w", err)
			}
			fmt.Println("\n[gitx main] Switched to 'main' and pulled all remotes.")
			return nil
		},
	}
}
