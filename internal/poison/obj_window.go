// @Title
// @Description
// @Author
// @Update
package poison

import (
	"fmt"
	"strings"

	"github.com/rivo/tview"
)

func (w *Window) CreateList(commmand func() []string, action func(string) func()) *tview.List {
	list := tview.NewList().
		ShowSecondaryText(false)

	status := commmand()
	for _, item := range status {
		if action == nil {
			list.AddItem(item, "", 0, nil)
		} else {
			list.AddItem(item, "", 0, action(item))
		}
	}

	return list
}

func (w *Window) CreateView(command func() []string) *tview.TextView {
	text := tview.NewTextView().
		SetTextAlign(tview.AlignLeft).
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true)

	results := command()
	for _, result := range results {
		if result != "" {
			fmt.Fprintf(text, "Remote : %s %s", tview.Escape(result), strings.Repeat(" ", 2))
		}
	}

	return text
}
