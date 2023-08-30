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
				if _, err := w.buttonModalSourceControl(buttonLabel); err != nil {
					w.App.Stop() // is temporary action
				}
			})
		w.Pages.AddPage("modal", modal, true, true)
		w.Pages.ShowPage("modal")
	}
}
