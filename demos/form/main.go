// Demo code for the Form primitive.
package main

import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	form := tview.NewForm().
		AddDropDown("Title", []string{"Mr.", "Ms.", "Mrs.", "Dr.", "Prof."}, 0, nil).
		AddInputField("First name", "", 20, nil).
		AddInputField("Last name", "", 20, nil).
		AddCheckbox("Age 18+", false, nil).
		AddButton("Save", nil).
		AddButton("Quit", func() {
			app.Stop()
		})
	form.SetBorder(true).SetTitle("Enter some data").SetTitleAlign(tview.AlignLeft)
	if err := app.SetRoot(form, true).SetFocus(form).Run(); err != nil {
		panic(err)
	}
}