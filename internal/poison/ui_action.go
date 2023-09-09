// @Title
// @Description
// @Author
// @Update
package poison

func (w *Window) buttonModalSourceControl(buttonLabel, item string) (str string, err error) {
	switch buttonLabel {
	case "Cancel":
		w.Pages.ShowPage("page1")
		w.Pages.HidePage("modal")
	case "Add":
		cmd := cmdGitAddItem(item)
		w.createModalAddItem(cmd)
		w.Pages.HidePage("modal")
	case "Discard":
	case "Diff":
	}
	return
}

func (w *Window) buttonModalAddItem(buttonLabel string) (str string, err error) {
	switch buttonLabel {
	case "Oke":
		w.LoadData()
		w.Pages.ShowPage("page1")
		w.Pages.HidePage("modalAddItem")
	}
	return
}
