package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// runCommand executes a shell command in the given directory
func runCommand(dir, cmd string, args ...string) {
	command := exec.Command(cmd, args...)
	command.Dir = dir
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	if err := command.Run(); err != nil {
		fmt.Println("Error:", err)
	}
}

func main() {
	// Ensure a directory is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: git-auto <project-path> [commit-message]")
		return
	}

	projectPath, err := filepath.Abs(os.Args[1]) // Get absolute path of project
	if err != nil {
		fmt.Println("Invalid path:", err)
		return
	}

	commitMessage := "Automated commit"
	if len(os.Args) > 2 {
		commitMessage = os.Args[2]
	}

	fmt.Println("Running Git automation in:", projectPath)

	// Run Git commands in the provided directory
	runCommand(projectPath, "git", "pull")
	runCommand(projectPath, "git", "add", ".")
	runCommand(projectPath, "git", "commit", "-m", commitMessage)
	runCommand(projectPath, "git", "push")
}
