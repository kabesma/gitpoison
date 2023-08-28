// @Title
// @Description
// @Author
// @Update
package poison

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (w *Window) CreateStatusView() *tview.List {
	list := tview.NewList().
		ShowSecondaryText(false).
		SetSecondaryTextColor(tcell.ColorDimGray)

	status := cmdGitStatus()
	for _, item := range status {
		list.AddItem(item, "", 0, w.createModalFunc(item))
	}

	return list
}

func (w *Window) createModalFunc(item string) func() {
	return func() {
		modal := tview.NewModal().
			SetText(fmt.Sprintf("Selected Action for this file :\n %s", item)).
			AddButtons([]string{"Add", "Discard", "Diff", "Cancel"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				if _, err := w.buttonModalStatus(buttonLabel); err != nil {
					w.App.Stop() // is temporary action
				}
			})
		w.Pages.AddPage("modal", modal, true, true)
		w.Pages.ShowPage("modal")
	}
}
