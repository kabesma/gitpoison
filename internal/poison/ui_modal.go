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

var (
	str string
)

func (w *Window) createModalSourceControl(item string) func() {
	return func() {
		if strings.HasPrefix(item, "(Staged)") {
			w.createModalSelected(item, []string{"Cancel", "Restore", "Diff"})
		} else {
			w.createModalSelected(item, []string{"Add", "Discard", "Diff", "Cancel"})
		}
	}
}

func (w *Window) createModalSelected(item string, btn []string) {
	w.ModalBasicGit = tview.NewModal().
		SetText(fmt.Sprintf("Selected Action for this file :\n %s", item)).
		AddButtons(btn).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if _, err := w.buttonModalSourceControl(buttonLabel, item); err != nil {
				w.App.Stop() // is temporary action
			}
		})
	w.Pages.AddPage("modalBasicGit", w.ModalBasicGit, true, true)
	w.Pages.ShowPage("modalBasicGit")
}

func (w *Window) createModalOk(item string) {
	w.ModalOk = tview.NewModal().
		SetText(fmt.Sprintf("%s", item)).
		AddButtons([]string{"Oke"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if _, err := w.buttonModalAddItem(buttonLabel); err != nil {
				w.App.Stop() // is temporary action
			}
		})
	w.Pages.AddPage("modalOk", w.ModalOk, true, true)
	w.Pages.ShowPage("modalOk")
}
