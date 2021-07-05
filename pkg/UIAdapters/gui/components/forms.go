package components

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/widget"
)

func SknSideForm() fyne.CanvasObject {
	name := widget.NewEntry()
	name.SetPlaceHolder("Network message")

	email := widget.NewEntry()
	email.SetPlaceHolder("test@example.com")
	email.Validator = validation.NewRegexp(`\w{1,}@\w{1,}\.\w{1,4}`, "not a valid email")

	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password")

	largeText := widget.NewMultiLineEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Name", Widget: name, HintText: "Your full name"},
			{Text: "Email", Widget: email, HintText: "A valid email address"},
		},
		OnCancel: func() {
			fmt.Println("Cancelled")
		},
		OnSubmit: func() {
			fmt.Println("Form submitted")
			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "Network message for: " + name.Text,
				Content: largeText.Text,
			})
		},
	}
	form.Append("Password", password)
	form.Append("Message", largeText)
	return form
}

func SknFirmwareSideForm() fyne.CanvasObject {
	name := widget.NewEntry()
	name.SetPlaceHolder("Firmware to load")

	email := widget.NewEntry()
	email.SetPlaceHolder("test@example.com")
	email.Validator = validation.NewRegexp(`\w{1,}@\w{1,}\.\w{1,4}`, "not a valid email")

	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password")

	largeText := widget.NewMultiLineEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Name", Widget: name, HintText: "URI file://, http://,"},
			{Text: "Email", Widget: email, HintText: "A valid email address"},
		},
		OnCancel: func() {
			fmt.Println("Cancelled")
		},
		OnSubmit: func() {
			fmt.Println("FirmWare submitted")
			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "Firmware for: " + name.Text,
				Content: largeText.Text,
			})
		},
	}
	form.Append("Password", password)
	form.Append("Message", largeText)
	return form
}

func SknScheduleSideForm() fyne.CanvasObject {
	name := widget.NewEntry()
	name.SetPlaceHolder("Device to Schedule")

	email := widget.NewEntry()
	email.SetPlaceHolder("test@example.com")
	email.Validator = validation.NewRegexp(`\w{1,}@\w{1,}\.\w{1,4}`, "not a valid email")

	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password")

	largeText := widget.NewMultiLineEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Name", Widget: name, HintText: "Your full name"},
			{Text: "Email", Widget: email, HintText: "A valid email address"},
		},
		OnCancel: func() {
			fmt.Println("Cancelled")
		},
		OnSubmit: func() {
			fmt.Println("Schedule submitted")
			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "Schedule for: " + name.Text,
				Content: largeText.Text,
			})
		},
	}
	form.Append("Password", password)
	form.Append("Message", largeText)
	return form
}
