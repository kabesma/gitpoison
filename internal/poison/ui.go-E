// @Title
// @Description
// @Author
// @Update
package poison

import (
	"kabesma/gitpoison/internal/component"

	"github.com/rivo/tview"
)

type Window struct {
	// Main Layout
	App           *tview.Application
	Pages         *tview.Pages
	Grid          *tview.Grid
	ModalBasicGit *tview.Modal
	ModalOk       *tview.Modal

	// Frame Screen
	SourceControl *tview.List
	Commits       *tview.List
	Branches      *tview.List
	Stashes       *tview.List
	Content       *tview.List
	// Content       *tview.TextView
	Status *tview.TextView

	// Commit
	Message *tview.InputField

	// Custom Modal
	ModalInput *component.CustomModal

	// Layer Page
	MainPage *tview.Flex

	// String
	BranchNow string

	// Mode Focus
	FocusMode bool
}
