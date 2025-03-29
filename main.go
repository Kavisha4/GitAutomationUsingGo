package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: git-auto <project-path> <command> [arguments]")
		return
	}

	projectPath := os.Args[1]
	command := os.Args[2]
	args := os.Args[3:]

	err := os.Chdir(projectPath)
	if err != nil {
		fmt.Println("Error: Unable to navigate to", projectPath)
		return
	}

	switch command {
	case "commit":
		runCommand("git", "add", ".")
		runCommand("git", "commit", "-m", args[0])
		runCommand("git", "pull", "--rebase")
		runCommand("git", "push")
	case "stash":
		runCommand("git", "stash")
	case "checkout":
		runCommand("git", "checkout", args[0])
	case "rebase":
		runCommand("git", "rebase", args[0])
	default:
		fmt.Println("Invalid command. Supported commands: commit, stash, checkout, rebase")
	}
}

func runCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(string(output))
}
