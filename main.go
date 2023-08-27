package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/rivo/tview"
)

type lane struct {
	pages *tview.Pages
}

func main() {
	l := lane{}
	app := tview.NewApplication()
	grid := tview.NewGrid()

	changes := l.CreateStatusView()
	changes.SetTitle("CHANGES").SetBorder(true)
	commit := tview.NewList().SetTitle("COMMIT").SetBorder(true)
	log := tview.NewTextView().SetTitle("LOG").SetBorder(true)

	sidebar := tview.NewGrid().SetRows(0, 0).
		AddItem(changes, 0, 0, 1, 1, 0, 0, true).
		AddItem(commit, 1, 0, 1, 1, 0, 0, false)

	content := tview.NewGrid().SetRows(0, 3).
		AddItem(log, 0, 0, 1, 1, 0, 0, false)

	grid = tview.NewGrid().SetRows(0, 2).
		SetColumns(70, 0).
		SetBorders(false).
		AddItem(sidebar, 0, 0, 1, 1, 0, 0, true).
		AddItem(content, 0, 1, 1, 1, 0, 0, true)

	// Create the first page
	page1 := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(grid, 0, 1, true)
		// AddItem(tview.NewBox().SetBorder(true).SetTitle("Page 1 - Box 2"), 0, 1, true)

	// Create the second page
	page2 := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Page 2 - Box 1"), 0, 1, true).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Page 2 - Box 2"), 0, 1, true)

	// Create a Pages layout
	l.pages = tview.NewPages().
		AddPage("page1", page1, true, true).
		AddPage("page2", page2, true, false)

	// Set the main layout and run the application
	if err := app.SetRoot(l.pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func GetGitStatus() []string {
	cmd := exec.Command("git", "status", "--short")

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

func CreateLogView() *tview.TextView {
	logView := tview.NewTextView().
		SetTextAlign(tview.AlignLeft).
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true)

	logs := GetGitStatus()
	for _, log := range logs {
		fmt.Fprintf(logView, "[%s]%s\n", tview.Escape(log), strings.Repeat(" ", 2))
	}

	return logView
}

func (l *lane) CreateStatusView() *tview.List {
	list := tview.NewList().ShowSecondaryText(false)

	status := GetGitStatus()
	for _, item := range status {
		list.AddItem(item, "", 0, l.createModalFunc(item))
	}

	return list
}

func colorizeStatusEntry(statusEntry string) string {
	parts := strings.SplitN(statusEntry, " ", 2)
	if len(parts) == 2 {
		return "[green]" + parts[0] + "[-] " + parts[1]
	}

	return statusEntry
}

func (l *lane) createModalFunc(item string) func() {
	return func() {
		modal := tview.NewModal().
			SetText(fmt.Sprintf("Selected Action for this file :\n %s", item)).
			AddButtons([]string{"Add", "Discard", "Diff", "Cancel"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				switch buttonLabel {
				case "Cancel":
					l.pages.ShowPage("page1")
					l.pages.HidePage("modal")
				case "Add":
				case "Discard":
				case "Diff":
				}
			})
		l.pages.AddPage("modal", modal, true, true)
		l.pages.ShowPage("modal")
	}
}
