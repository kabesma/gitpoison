// @Key Event
// @This for control all Key
// @ak4bento
// @ak4bento
package poison

import (
	"sync"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (w *Window) setupKeyboard() {
	focusMapping := map[tview.Primitive]struct{ next, prev tview.Primitive }{
		w.SourceControl: {w.Commits, w.Status},
		w.Commits:       {w.Branches, w.SourceControl},
		w.Branches:      {w.Content, w.Commits},
		w.Content:       {w.Stashes, w.Branches},
		w.Stashes:       {w.Status, w.Content},
		w.Status:        {w.SourceControl, w.Stashes},
	}

	// Setup app level keyboard shortcuts.
	w.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case KeyMapping[KeySourcesControl]:
			w.App.SetFocus(w.SourceControl)
		case KeyMapping[KeyCommits]:
			w.setFocus(w.Commits)
		case KeyMapping[KeyBranches]:
			w.setFocus(w.Branches)
		case KeyMapping[KeyContent]:
			w.setFocus(w.Content)
		case KeyMapping[KeyStashes]:
			w.setFocus(w.Stashes)
		case KeyMapping[KeyStatus]:
			w.setFocus(w.Status)
		case tcell.KeyCtrlO:
			w.createModalSelected("All", []string{"Add", "Discard", "Diff", "Cancel"})
		case tcell.KeyCtrlR:
			w.LoadData()
		case tcell.KeyCtrlP:
			w.createModalCommit()
		case tcell.KeyCtrlF:
			w.createModalConfirm(func() {
				var wg sync.WaitGroup
				wg.Add(1)
				message := w.ModalInput.InputField.GetText()
				cmdGitCommit(message)
				aaaa, err := cmdGitPush(w.BranchNow, &wg)
				if err != nil {
					w.createModalOk(err.Error())
				}
				w.createModalOk("Successfully executed\n" + aaaa)
				w.Pages.HidePage("modalConfirm")

			})
		// w.toggleFocusMode()
		// case tcell.KeyEscape:
		// if w.FocusMode {
		// w.toggleFocusMode()
		// }

		/* Configuration for Tab & Backtab keys */

		// On Tab press set focus to the next element.
		case tcell.KeyTab:
			if focusMap, ok := focusMapping[w.App.GetFocus()]; ok {
				w.setFocus(focusMap.next)
			}

			// Return `nil` to avoid default Backtab behaviour for the primitive.
			return nil

		// On Backtab press set focus to the prev element.
		case tcell.KeyBacktab:
			if focusMap, ok := focusMapping[w.App.GetFocus()]; ok {
				w.setFocus(focusMap.prev)
			}

			// Return `nil` to avoid default Backtab behaviour for the primitive.
			return nil
		}
		return event
	})

	// Setup Tables element level keyboard shortcuts.
	// w.Tables.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
	// switch event.Rune() {
	// case 'e':
	// w.describeSelectedTable()
	// case 'p':
	// w.previewSelectedTable()
	// }
	// return event
	// })
}

func (w *Window) handlerCommit(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyEnter:
		w.Pages.ShowPage("page1")
		w.Pages.HidePage("modalCommit")
		message := w.ModalInput.InputField.GetText()
		strMessage, err := cmdGitCommit(message)
		if err != nil {
			return event
		}
		w.createModalOk(strMessage)
	case tcell.KeyCtrlK:
		var isDone bool

		w.createModalConfirm(func() {
			var wg sync.WaitGroup
			wg.Add(1)
			message := w.ModalInput.InputField.GetText()
			cmdGitCommit(message)

			go cmdGitPush(w.BranchNow, &wg)
			wg.Wait()
			go func() {
				wg.Wait() // Tunggu sampai semua goroutine selesai
				isDone = true
			}()

			for !isDone {
				// Tunggu sampai isDone menjadi true
			}
			w.createModalOk("Successfully executed")
			w.Pages.HidePage("modalConfirm")
		})
		w.Pages.HidePage("modalCommit")
	case tcell.KeyEscape:
		w.Pages.ShowPage("page1")
		w.Pages.HidePage("modalCommit")
		w.LoadData()
	}

	return event
}

func (w *Window) setFocus(p tview.Primitive) {
	w.queueUpdateDraw(func() {
		w.App.SetFocus(p)
	})
}

func (w *Window) queueUpdateDraw(f func()) {
	go func() {
		w.App.QueueUpdateDraw(f)
	}()
}
