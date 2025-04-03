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
  "bytes"
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
		status[i] = themeColorSourceControl(item)
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
		status[i] = themeColorBranch(item)
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

func cmdGitBranchCurrent() string {
	cmd := exec.Command("git", "branch", "--show-current")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing 'git branch' command: %s\n", err)
		return ""
	}

	return string(output)
}

func cmdGitCommit(message string) (string, error) {
	cmd := exec.Command("git", "commit", "-m", message)

	_, err := cmd.CombinedOutput()
	if err != nil {
		return "Error executing 'git commit'\n command : " + err.Error(), err
	}

	return "Successfully executed", nil
}

// Fungsi untuk mendapatkan path repository Git saat ini
func getRepoPath() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("Gagal mendapatkan repo path: %v", err)
	}
	return strings.TrimSpace(out.String()), nil
}

// Fungsi untuk melakukan git push
func cmdGitPush(branch string) (string, error) {
	// Ambil repo path secara dinamis
	repoPath, err := getRepoPath()
	if err != nil {
		return fmt.Sprintf("Error: %v", err), err
	}

	// Buat command git push
	cmd := exec.Command("git", "push", "origin", branch)
	cmd.Dir = repoPath // Set working directory ke repo Git

	// Jalankan command dan tangkap output
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("Error executing 'git push': %s\nOutput: %s", err.Error(), string(output)), err
	}

	return "Successfully executed", nil
}
// func cmdGitPush(branch string) (string, error) {
// 	cmd := exec.Command("git", "push", "origin", branch)
//
// 	_, err := cmd.CombinedOutput()
// 	if err != nil {
// 		return "Error executing 'git push'\n command : " + err.Error(), err
// 	}
//
// 	return "Successfully executed", nil
// }

func cmdGitAddItem(item string) string {
	if item == "All" {
		cmd := exec.Command("git", "add", ".")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return "Error executing 'git add .'\n command : " + err.Error()
		}
		return "You have added " + item + "\n" + string(output)
	}

	split := strings.SplitN(item, " ", 2)
	if len(split) == 2 {
		cmd := exec.Command("git", "add", split[1])
		output, err := cmd.CombinedOutput()
		if err != nil {
			return "Error executing 'git add'\n command : " + err.Error()
		}
		return "You have added " + item + "\n" + string(output)
	}

	return "Error executing 'git add'"
}

func cmdGitRestoreStaged(item string) string {
	if item == "All" {
		cmd := exec.Command("git", "restore", "--staged", ".")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return "Error executing 'git restore .'\n command : " + err.Error()
		}
		return "You have restored " + item + "\n" + string(output)
	}

	split := strings.SplitN(item, " ", 2)
	if len(split) == 2 {
		cmd := exec.Command("git", "restore", "--staged", split[1])
		output, err := cmd.CombinedOutput()
		if err != nil {
			return "Error executing 'git restore --staged " + split[1] + "'\n command : " + err.Error()
		}
		return "You have restored " + item + "\n" + string(output)
	}

	return "Error executing 'git restore'"
}

func cmdGitRestoreChanged(item string) string {
	if item == "All" {
		cmd := exec.Command("git", "restore", ".")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return "Error executing 'git restore .'\n command : " + err.Error()
		}
		return "You have restored " + item + "\n" + string(output)
	}

	split := strings.SplitN(item, " ", 2)
	if len(split) == 2 {
		cmd := exec.Command("git", "restore", split[1])
		output, err := cmd.CombinedOutput()
		if err != nil {
			return "Error executing 'git restore " + split[1] + "'\n command : " + err.Error()
		}
		return "You have restored " + item + "\n" + string(output)
	}

	return "Error executing 'git restore'"
}
