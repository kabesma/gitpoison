// @Title
// @Description
// @Author
// @Update
package poison

func (w *Window) buttonModalSourceControl(buttonLabel, item string) (str string, err error) {
	switch buttonLabel {
	case "Cancel":
		w.LoadData()
		w.Pages.ShowPage("page1")
		w.Pages.HidePage("modalBasicGit")
	case "Add":
		cmd := cmdGitAddItem(item)
		w.createModalOk(cmd)
		w.Pages.HidePage("modalBasicGit")
	case "Restore":
		restore := cmdGitRestoreStaged(item)
		w.createModalOk(restore)
		w.Pages.HidePage("modalBasicGit")
	case "Discard":
		restore := cmdGitRestoreChanged(item)
		w.createModalOk(restore)
		w.Pages.HidePage("modalBasicGit")
	case "Diff":
	}
	return
}

func (w *Window) buttonModalAddItem(buttonLabel string) (str string, err error) {
	switch buttonLabel {
	case "Oke":
		w.LoadData()
		w.Pages.ShowPage("page1")
		w.Pages.HidePage("modalOk")
	}
	return
}
