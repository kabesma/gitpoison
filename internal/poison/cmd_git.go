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
