package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// Helper function to run shell commands
func runCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

// Root command
var rootCmd = &cobra.Command{
	Use:   "git-auto",
	Short: "Automates Git workflows",
}

// Commit command
var commitCmd = &cobra.Command{
	Use:   "commit [message]",
	Short: "Commit changes with a message",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runCommand("git", "add", ".")
		runCommand("git", "commit", "-m", args[0])
		runCommand("git", "push")
	},
}

// Stash command
var stashCmd = &cobra.Command{
	Use:   "stash",
	Short: "Stash changes",
	Run: func(cmd *cobra.Command, args []string) {
		runCommand("git", "stash")
	},
}

// Checkout command
var checkoutCmd = &cobra.Command{
	Use:   "checkout [branch]",
	Short: "Switch branches",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runCommand("git", "checkout", args[0])
	},
}

// Rebase command
var rebaseCmd = &cobra.Command{
	Use:   "rebase [branch]",
	Short: "Rebase current branch onto another",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runCommand("git", "rebase", args[0])
	},
}

func main() {
	rootCmd.AddCommand(commitCmd)
	rootCmd.AddCommand(stashCmd)
	rootCmd.AddCommand(checkoutCmd)
	rootCmd.AddCommand(rebaseCmd)
	rootCmd.Execute()
}
