// @Title
// @Description
// @Author
// @Update
package poison

import (
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
		// case tcell.KeyCtrlR:
		// w.LoadData()
		// case tcell.KeyCtrlF:
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