// @Title
// @Description
// @Author
// @Update
package poison

import "strings"

func colorizeStatusEntry(statusEntry string) string {
	parts := strings.SplitN(statusEntry, " ", 2)
	if len(parts) == 2 {
		statusSymbol := parts[0]
		if parts[0] == "??" {
			statusSymbol = "[green]U"
			return statusSymbol + "[-] " + parts[1]
		}

		if res := strings.SplitN(parts[1], " ", 2); len(res) == 2 {
			switch res[0] {
			case "M":
				statusSymbol = "[yellow]M"
			case "D":
				statusSymbol = "[red]D"
			}
			return statusSymbol + "[-] " + res[1]
		}
	}

	return "All"
}
