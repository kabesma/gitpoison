// @Title
// @Description
// @Author
// @Update
package poison

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func cmdGitLogCommit() []string {
	cmd := exec.Command("git", "log", "--pretty=format:%s - %ar")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing 'git log' command: %s\n", err)
		return []string{}
	}

	log := strings.Split(string(output), "\n")
	for i, item := range log {
		log[i] = item
	}

	return log
}

func cmdGitStatus() []string {
	cmd := exec.Command("git", "status", "--short", "-u")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing 'git status' command: %s\n", err)
		return []string{}
	}

	status := strings.Split(string(output), "\n")
	for i, item := range status {
		status[i] = colorizeStatusEntry(item)
	}

	return status
}

func cmdGitBranch() []string {
	cmd := exec.Command("git", "branch", "--color=auto")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing 'git branch' command: %s\n", err)
		return []string{}
	}

	status := strings.Split(string(output), "\n")
	for i, item := range status {
		status[i] = colorizeBranchEntry(item)
	}

	return status
}

func cmdGitStash() []string {
	cmd := exec.Command("git", "stash", "list")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing 'git stash' command: %s\n", err)
		return []string{}
	}

	status := strings.Split(string(output), "\n")
	for i, item := range status {
		status[i] = item
	}

	return status
}

func cmdGitLogGraph() []string {
	cmd := exec.Command(
		"git",
		"log",
		"--all",
		"--decorate",
		"--graph", "--pretty=format:%h (%an) %s - %ar %b")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing 'git log graph' command: %s\n", err)
		return []string{}
	}

	status := strings.Split(string(output), "\n")
	for i, item := range status {
		status[i] = item
	}

	return status
}

func cmdGitBranchCurrent() []string {
	cmd := exec.Command("git", "branch", "--show-current")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing 'git branch' command: %s\n", err)
		return []string{}
	}

	status := strings.Split(string(output), "\n")
	for i, item := range status {
		status[i] = item
	}

	return status
}
