package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/user/gitx/internal/reclone"
)

var (
	Version = "0.0.0-dev"
)

var rootCmd = &cobra.Command{
	Use:     "gitx",
	Short:   "gitx is a family of x-suffix utilities for git",
	Long:    `gitx is a Go-based command-line tool that provides enhanced git workflows.`,
	Version: Version,
}

func init() {
	rootCmd.AddCommand(reclone.NewRecloneCmd())
	rootCmd.AddCommand(NewMasterCmd())
	rootCmd.AddCommand(NewMainCmd())
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
