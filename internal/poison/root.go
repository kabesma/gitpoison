// @Title
// @Description
// @Author
// @Update
package poison

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Execute() {
	w := Window{}

	w.App = tview.NewApplication()
	w.Grid = tview.NewGrid()

	w.SourceControl = w.CreateStatusView()
	// tview.NewList().
	// ShowSecondaryText(true).
	// SetSecondaryTextColor(tcell.ColorDimGray)
	w.Commits = tview.NewList().
		ShowSecondaryText(false)
	w.Branches = tview.NewList().
		ShowSecondaryText(false)
	w.Stashes = tview.NewList().
		ShowSecondaryText(false)
	w.Content = tview.NewTextView().
		SetTextAlign(tview.AlignLeft).
		SetTextColor(tcell.ColorGray)
	w.Status = tview.NewTextView().
		SetTextAlign(tview.AlignLeft).
		SetTextColor(tcell.ColorGray)

	w.SourceControl.SetTitle("SOURCE CONTROL").SetBorder(true)
	w.Commits.SetTitle("COMMITS").SetBorder(true)
	w.Branches.SetTitle("BRANCHES").SetBorder(true)
	w.Stashes.SetTitle("STASHES").SetBorder(true)
	w.Content.SetTitle("RESULT").SetBorder(true)
	w.Status.SetTitle("STATUS").SetBorder(true)

	sidebar := tview.NewGrid().
		SetRows(0, 0, 0, 0).
		AddItem(w.SourceControl, 0, 0, 1, 1, 0, 0, true).
		AddItem(w.Commits, 1, 0, 1, 1, 0, 0, false).
		AddItem(w.Branches, 2, 0, 1, 1, 0, 0, false).
		AddItem(w.Stashes, 3, 0, 1, 1, 0, 0, false)

	content := tview.NewGrid().
		SetRows(0, 5).
		AddItem(w.Content, 0, 0, 1, 1, 0, 0, false).
		AddItem(w.Status, 1, 0, 1, 1, 0, 0, false)

	w.Grid = tview.NewGrid().
		SetRows(0).
		SetColumns(50, 0).
		SetBorders(false).
		AddItem(sidebar, 0, 0, 1, 1, 0, 0, true).
		AddItem(content, 0, 1, 1, 1, 0, 0, true)

	w.MainPage = tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(w.Grid, 0, 1, true)

	w.Pages = tview.NewPages().
		AddPage("MainPage", w.MainPage, true, true)

	if err := w.App.SetRoot(w.Pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
