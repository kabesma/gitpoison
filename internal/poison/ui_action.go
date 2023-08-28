// @Title
// @Description
// @Author
// @Update
package poison

func (w *Window) buttonModalStatus(buttonLabel string) (str string, err error) {
	switch buttonLabel {
	case "Cancel":
		w.Pages.ShowPage("page1")
		w.Pages.HidePage("modal")
	case "Add":
	case "Discard":
	case "Diff":
	}
	return
}
