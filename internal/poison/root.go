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

type (
	// KeyOp defines dbui specific hotkey operations.
	KeyOp int16
)

const (
	// KeySourcesOp is the operation corresponding to the activation of the Sources view.
	KeySourcesControl KeyOp = iota
	// KeySchemasOp is the operation corresponding to the activation of the Schemas view.
	KeyCommits
	// KeyTablesOp is the operation corresponding to the activation of the Tables view.
	KeyBranches
	// KeyPreviewOp is the operation corresponding to the activation of the Preview view.
	KeyStashes
	// KeyQyOp is the operation corresponding to the activation of the Query view.
	KeyContent
	KeyStatus
)

var (
	// KeyMapping maps keyboard operations to their hotkeys. In the future, this part can be customized by the user configuration.
	KeyMapping = map[KeyOp]tcell.Key{
		KeySourcesControl: tcell.KeyCtrlA,
		KeyCommits:        tcell.KeyCtrlS,
		KeyBranches:       tcell.KeyCtrlD,
		KeyStashes:        tcell.KeyCtrlE,
		KeyContent:        tcell.KeyCtrlQ,
		KeyStatus:         tcell.KeyCtrlW,
	}

	TitleSourcesControl = fmt.Sprintf("SOURCE CONTROL [ %s ]", tcell.KeyNames[KeyMapping[KeySourcesControl]])
	TitleCommits        = fmt.Sprintf("COMMITS [ %s ]", tcell.KeyNames[KeyMapping[KeyCommits]])
	TitleBranches       = fmt.Sprintf("BRANCHES [ %s ]", tcell.KeyNames[KeyMapping[KeyBranches]])
	TitleStashes        = fmt.Sprintf("STASHES [ %s ]", tcell.KeyNames[KeyMapping[KeyStashes]])
	TitleContent        = fmt.Sprintf("RESULT [ %s ]", tcell.KeyNames[KeyMapping[KeyContent]])
	TitleStatus         = fmt.Sprintf("STATUS [ %s ]", tcell.KeyNames[KeyMapping[KeyStatus]])
)

func Execute() {
	w := Window{}

	w.App = tview.NewApplication()
	w.Grid = tview.NewGrid()

	w.SourceControl = w.CreateList(cmdGitStatus, w.createModalSourceControl)
	w.Commits = w.CreateList(cmdGitLogCommit, nil)
	w.Branches = w.CreateList(cmdGitBranch, nil)
	w.Stashes = w.CreateList(cmdGitStash, nil)

	w.Content = w.CreateList(cmdGitLogGraph, nil)

	w.Status = w.CreateView(cmdGitBranchCurrent)
	// tview.NewTextView().
	// SetTextAlign(tview.AlignLeft).
	// SetTextColor(tcell.ColorGray)

	w.SourceControl.SetTitle(TitleSourcesControl).SetBorder(true)
	w.Commits.SetTitle(TitleCommits).SetBorder(true)
	w.Branches.SetTitle(TitleBranches).SetBorder(true)
	w.Stashes.SetTitle(TitleStashes).SetBorder(true)
	w.Content.SetTitle(TitleContent).SetBorder(true)
	w.Status.SetTitle(TitleStatus).SetBorder(true)

	sidebar := tview.NewGrid().
		SetRows(3, 0, 0, 0).
		AddItem(w.Status, 0, 0, 1, 1, 0, 0, true).
		AddItem(w.SourceControl, 1, 0, 1, 1, 0, 0, true).
		AddItem(w.Commits, 2, 0, 1, 1, 0, 0, false).
		AddItem(w.Branches, 3, 0, 1, 1, 0, 0, false)

	content := tview.NewGrid().
		SetRows(0, 10).
		AddItem(w.Content, 0, 0, 1, 1, 0, 0, false).
		AddItem(w.Stashes, 1, 0, 1, 1, 0, 0, false)

	w.Grid = tview.NewGrid().
		SetRows(0).
		SetColumns(50, 0).
		SetBorders(false).
		AddItem(sidebar, 0, 0, 1, 1, 0, 0, true).
		AddItem(content, 0, 1, 1, 1, 0, 0, false)

	w.MainPage = tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(w.Grid, 0, 1, true)

	w.setupKeyboard()
	w.Pages = tview.NewPages().
		AddPage("MainPage", w.MainPage, true, true)

	if err := w.App.SetRoot(w.Pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
