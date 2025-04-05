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

func (w *Window) CreateView(command func() string) *tview.TextView {
	text := tview.NewTextView().
		SetTextAlign(tview.AlignLeft).
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true)

	result := command()
	w.BranchNow = result
	if result != "" {
		fmt.Fprintf(text, "Remote : %s %s", tview.Escape(result), strings.Repeat(" ", 2))
	}

	return text
}
