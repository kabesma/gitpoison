// @Title
// @Description
// @Author
// @Update
package poison

import (
	"strings"
)

func selectionStatusStagged(str string, first, last int) (changed string) {
	changed = str
	switch str[first:last] {
	case "?":
		changed = "[blue]U"
	case "M":
		changed = "[blue]M"
	case "D":
		changed = "[blue]D"
	case "A":
		changed = "[blue]A"
	}
	return
}
func selectionStatusChanged(str string, first, last int) (changed string) {
	changed = str
	switch str[first:last] {
	case "?":
		changed = "[green]U"
	case "M":
		changed = "[yellow]M"
	case "D":
		changed = "[red]D"
	case "A":
		changed = "[green]A"
	}
	return
}

func colorizeStatusEntry(statusEntry string) string {
	parts := strings.SplitN(statusEntry, " ", 2)
	if len(parts) == 2 {
		// result := strings.ReplaceAll(parts[0], " ", "")
		changed := ""
		stagged := ""
		info := ""

		if parts[0] != "" {
			countIndex := len(parts[0])
			result := parts[0]

			if countIndex == 1 {
				changed = selectionStatusChanged(result, 0, countIndex)
			} else if countIndex == 2 {
				changed = selectionStatusChanged(result, 0, countIndex-1)
				if result[0:1] != "?" {
					stagged = selectionStatusStagged(result, 1, countIndex)
					info = "(Stagged) "
				}
			}
			// return result[0:countIndex] + " " + strconv.Itoa(countIndex)
			return info + changed + stagged + "[-] " + parts[1] //+ strconv.Itoa(countIndex)
		}

		if res := strings.SplitN(parts[1], " ", 2); len(res) == 2 {
			countIndex := len(res[0])
			result := res[0]

			if countIndex == 1 {
				changed = selectionStatusChanged(result, 0, countIndex)
			} else if countIndex == 2 {
				changed = selectionStatusChanged(result, 0, countIndex-1)
				if result[0:1] != "?" {
					stagged = selectionStatusStagged(result, 1, countIndex)
					info = "(Stagged) "
				}
			}
			// return result[0:countIndex] + " " + strconv.Itoa(countIndex)
			return info + changed + stagged + "[-] " + res[1]
		}
		// statusSymbol := parts[0]
		// if parts[0] == "??" {
		// statusSymbol = "[green]U"
		// return statusSymbol + "[-] " + parts[1]
		// }
		// if parts[0] != "" {
		// switch parts[0] {
		// case "??":
		// statusSymbol = "[green]U"
		// case "M":
		// statusSymbol = "[yellow]M"
		// case "D":
		// statusSymbol = "[red]D"
		// case "A":
		// statusSymbol = "[green]A"
		// }
		// return "(Stagged) " + statusSymbol + "[-] " + parts[1]
		// }
		//
		// if res := strings.SplitN(parts[1], " ", 2); len(res) == 2 {
		// switch res[0] {
		// case "M":
		// statusSymbol = "[yellow]M"
		// case "D":
		// statusSymbol = "[red]D"
		// }
		// return statusSymbol + "[-] " + res[1]
		// }
	}

	return "All"
}

func colorizeBranchEntry(branchEntry string) string {
	parts := strings.SplitN(branchEntry, " ", 2)
	if parts[0] == "*" {
		return parts[0] + "[-] [green]" + parts[1]
	}

	return branchEntry
}
