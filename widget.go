package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2/widget"
)

type Widgets struct {
	MainForm      *widget.Form
	FilenameText  *widget.Label
	FileSelectBtn *widget.Button
	Form          *widget.Form
	Timezone      *widget.Select
	TargetServer  *widget.SelectEntry
	Status        *widget.TextGrid
}

func (w *Widgets) SetWidgets(config *config) {
	w.Status = widget.NewTextGrid()
	label := widget.NewLabel(config.UploadFilename)
	label.SetText(config.UploadFilename)

	w.FilenameText = label

	timezone := widget.NewSelect([]string{"Asia/Seoul", "America/New_York"},
		func(tz string) {
			config.Timezone = tz
		})

	timezone.Selected = config.Timezone

	w.Timezone = timezone

	btn1 := widget.NewButton("open",
		func() {
			fmt.Println("open")
		},
	)

	w.FileSelectBtn = btn1

	targetServer := widget.NewSelectEntry([]string{"Dev Server", "Test Server", "Release Server"})
	targetServer.OnChanged = func(val string) {
		switch val {
		case "Dev Server":
			config.ServerAddress = "https://iotdevserver.inhandplus.com"
		case "Test Server":
			config.ServerAddress = "https://iottestserver.inhandplus.com"
		case "Release Server":
			config.ServerAddress = "https://iotserver.inhandplus.com"
		default:
			config.ServerAddress = val
		}
	}

	switch config.ServerAddress {
	case "https://iotdevserver.inhandplus.com":
		config.ServerAddress = "Dev Server"
	case "https://iottestserver.inhandplus.com":
		config.ServerAddress = "Test Server"
	case "https://iotserver.inhandplus.com":
		config.ServerAddress = "Release Server"
	default:
		break
	}
	targetServer.Text = config.ServerAddress

	targetServer.OnSubmitted = func(val string) {
		config.ServerAddress = val
	}

	w.TargetServer = targetServer

	mainForm := widget.NewForm(
		widget.NewFormItem("File name", w.FilenameText),
		widget.NewFormItem("File select button", btn1),
		widget.NewFormItem("Target server", w.TargetServer),
		widget.NewFormItem("Timezone", w.Timezone),
	)

	mainForm.OnSubmit = func() {
		err := uploading(config)

		if err != nil {
			sbox.AddLine("Upload failed.")
		}
	}

	mainForm.OnCancel = func() {
		os.Exit(0)
	}

	mainForm.SubmitText = "Upload"
	mainForm.CancelText = "Quit"

	w.MainForm = mainForm
}
