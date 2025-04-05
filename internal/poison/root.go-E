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
	KeySourcesControl KeyOp = iota
	KeyCommits
	KeyBranches
	KeyStashes
	KeyContent
	KeyStatus
)

var (
	// KeyMapping maps keyboard operations to their hotkeys. In the future, this part can be customized by the user configuration.
	KeyMapping = map[KeyOp]tcell.Key{
		KeySourcesControl: tcell.KeyCtrlX,
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

func Execute() *Window {
	w := Window{}

	w.App = tview.NewApplication()
	w.Grid = tview.NewGrid()

	w.SourceControl = tview.NewList().ShowSecondaryText(false)
	w.Commits = tview.NewList().ShowSecondaryText(false)
	w.Branches = tview.NewList().ShowSecondaryText(false)
	w.Stashes = tview.NewList().ShowSecondaryText(false)
	w.Content = tview.NewList().ShowSecondaryText(false)

	w.Status = w.CreateView(cmdGitBranchCurrent)

	w.Message = tview.NewInputField()
	// tview.NewTextView().
	// SetTextAlign(tview.AlignLeft).
	// SetTextColor(tcell.ColorGray)

	w.SourceControl.SetTitle(TitleSourcesControl).SetBorder(true)
	w.Commits.SetTitle(TitleCommits).SetBorder(true)
	w.Branches.SetTitle(TitleBranches).SetBorder(true)
	w.Stashes.SetTitle(TitleStashes).SetBorder(true)
	w.Content.SetTitle(TitleContent).SetBorder(true)
	w.Status.SetTitle(TitleStatus).SetBorder(true)
	w.Message.SetTitle("MESSAGE").SetBorder(true)

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
	w.LoadData()
	w.Pages = tview.NewPages().
		AddPage("MainPage", w.MainPage, true, true)

	return &w
}

func (w *Window) StartApp() {
	// w.App.Draw()
	if err := w.App.SetRoot(w.Pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func (w *Window) LoadData() {
	w.SourceControl.Clear()
	w.Commits.Clear()
	w.Branches.Clear()
	w.Stashes.Clear()
	w.Content.Clear()
	// w.Status.Clear()

	sourceControl := cmdGitStatus()
	commits := cmdGitLogCommit()
	branches := cmdGitBranch()
	stashes := cmdGitStash()
	content := cmdGitLogGraph()

	w.queueUpdate(func() {
		for _, item := range sourceControl {
			w.SourceControl.AddItem(item, "", 0, w.createModalSourceControl(item))
		}
		w.setFocus(w.SourceControl)

		for _, item := range commits {
			w.Commits.AddItem(item, "", 0, nil)
		}

		for _, item := range branches {
			w.Branches.AddItem(item, "", 0, nil)
		}

		for _, item := range stashes {
			w.Stashes.AddItem(item, "", 0, nil)
		}

		for _, item := range content {
			w.Content.AddItem(item, "", 0, nil)
		}
	})
}

func (w *Window) queueUpdate(f func()) {
	go func() {
		w.App.QueueUpdate(f)
	}()
}
