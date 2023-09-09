// @Title
// @Description
// @Author
// @Update
package poison

import (
	"fmt"

	"github.com/rivo/tview"
)

var (
	str string
)

func (w *Window) createModalSourceControl(item string) func() {
	return func() {
		modal := tview.NewModal().
			SetText(fmt.Sprintf("Selected Action for this file :\n %s", item)).
			AddButtons([]string{"Add", "Discard", "Diff", "Cancel"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				if _, err := w.buttonModalSourceControl(buttonLabel, item); err != nil {
					w.App.Stop() // is temporary action
				}
			})
		w.Pages.AddPage("modal", modal, true, true)
		w.Pages.ShowPage("modal")
	}
}

func (w *Window) createModalSelected(item string) {
	modal := tview.NewModal().
		SetText(fmt.Sprintf("Selected Action for this file :\n %s", item)).
		AddButtons([]string{"Add", "Discard", "Diff", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if _, err := w.buttonModalSourceControl(buttonLabel, item); err != nil {
				w.App.Stop() // is temporary action
			}
		})
	w.Pages.AddPage("modal", modal, true, true)
	w.Pages.ShowPage("modal")
}

func (w *Window) createModalAddItem(item string) {
	w.ModalAddItem = tview.NewModal().
		SetText(fmt.Sprintf("%s", item)).
		AddButtons([]string{"Oke"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if _, err := w.buttonModalAddItem(buttonLabel); err != nil {
				w.App.Stop() // is temporary action
			}
		})
	w.Pages.AddPage("modalAddItem", w.ModalAddItem, true, true)
	w.Pages.ShowPage("modalAddItem")
}
