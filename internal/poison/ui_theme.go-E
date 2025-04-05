// @Title
// @Description
// @Author
// @Update
package poison

import (
	"strings"
)

func themeColorSourceControl(entry string) string {
	if entry == "" {
		return "All"
	}

	firstTwoChars := entry[:2]
	restOfTheString := entry[2:]

	result := ""

	if strings.HasPrefix(firstTwoChars, " ") {
		result = "[gray]" + strings.ReplaceAll(firstTwoChars, " ", "")
	} else if firstTwoChars[:1] != " " {
		firstChar := themeSetConvertChars(firstTwoChars, "[green]", 0, 1)
		secondChar := themeSetConvertChars(firstTwoChars, "[red]", 1, 2)
		result = themeSetStaged(firstChar, secondChar)
	}

	return result + "[white] " + strings.ReplaceAll(restOfTheString, " ", "")
}

func themeSetStaged(firstChar, secondChar string) string {
	if secondChar == "" {
		return "(Staged)" + firstChar
	}
	return firstChar + secondChar
}

func themeSetConvertChars(str, color string, first, last int) string {
	chars := str
	switch chars[first:last] {
	case "?":
		// untracked
		chars = color + "U"
	case "!":
		// ignored
		chars = color + "I"
	case "M":
		// modified
		chars = color + "M"
	case "D":
		// deleted
		chars = color + "D"
	case "A":
		// added
		chars = color + "A"
	case "R":
		// renamed
		chars = color + "R"
	case "C":
		// copied
		chars = color + "C"
	case "T":
		// type changed
		chars = color + "T"
	case " ":
		// nothing
		chars = ""
	}
	return chars
}

func themeColorBranch(branchEntry string) string {
	parts := strings.SplitN(branchEntry, " ", 2)
	if parts[0] == "*" {
		return parts[0] + "[-] [green]" + parts[1]
	}

	return branchEntry
}
