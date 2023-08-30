// @Title
// @Description
// @Author
// @Update
package poison

import "github.com/rivo/tview"

type Window struct {
	// Main Layout
	App   *tview.Application
	Pages *tview.Pages
	Grid  *tview.Grid
	Modal *tview.Modal

	// Frame Screen
	SourceControl *tview.List
	Commits       *tview.List
	Branches      *tview.List
	Stashes       *tview.List
	Content       *tview.List
	// Content       *tview.TextView
	Status *tview.TextView

	// Layer Page
	MainPage *tview.Flex
}
