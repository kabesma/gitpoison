// @Title
// @Description
// @Author
// @Update
package component

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// CustomModal is a custom modal component that implements tview.Primitive.
type CustomModal struct {
	*tview.InputField

	frame *tview.Frame
}

// NewCustomModal creates a new instance of CustomModal with the specified width and height.
func NewCustomModal() *CustomModal {
	m := &CustomModal{
		InputField: tview.NewInputField(),
	}

	m.InputField.SetTitle("Message Your Commit").SetBorder(true)

	m.InputField.SetBorderColor(tview.Styles.PrimaryTextColor).
		SetBackgroundColor(tview.Styles.ContrastBackgroundColor)

	m.frame = tview.NewFrame(m.InputField).SetBorders(0, 0, 1, 0, 0, 0)
	m.frame.SetBorder(true).
		SetBackgroundColor(tview.Styles.ContrastBackgroundColor).
		SetBorderPadding(1, 1, 2, 1)
	return m
}

func (c *CustomModal) SetInputCapture(fn func(event *tcell.EventKey) *tcell.EventKey) {
	c.InputField.SetInputCapture(fn)
}

// Draw draws the modal.
func (c *CustomModal) Draw(screen tcell.Screen) {
	screenWidth, screenHeight := screen.Size()
	height := 0
	width := screenWidth / 3

	c.frame.Clear()

	c.frame.AddText("Enter : Commit", true, tview.AlignLeft, tview.Styles.PrimaryTextColor).
		AddText("Esc : Cancel", true, tview.AlignLeft, tview.Styles.PrimaryTextColor).
		AddText("Ctrl + K : Commit and Push", true, tview.AlignLeft, tview.Styles.PrimaryTextColor)

	// Set the modal's position and size.
	height += 4
	width += 4
	x := (screenWidth - width) / 2
	y := (screenHeight - height) / 2

	c.frame.SetRect(x, y-4, width+4, height+7)
	c.frame.Draw(screen)
	c.InputField.SetRect(x+2, y+1, width, height-5)
}
