package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var sbox *StatusBox

func makeUI() (*widget.Label, *widget.Entry) {
	out := widget.NewLabel("Hello world!")
	in := widget.NewEntry()

	in.OnChanged = func(content string) {
		out.SetText(("Hello " + content + "!"))
	}
	return out, in
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Video clip uploader")

	config := config{}

	sbox = NewStatusBox(8)

	sbox.AddLine("InHandPlus Video Clip Uploader v1.0\n")

	err := config.LoadConfig()

	if err != nil {
		sbox.AddLine(err.Error())
	} else {
		sbox.AddLine("Loaded your saved settings\n")
	}

	widgets := Widgets{MainWindow: myWindow}

	widgets.SetWidgets(&config)

	myWindow.SetContent(
		container.NewPadded(
			container.NewBorder(
				widgets.MainForm, widgets.Progress, container.NewPadded(), container.NewPadded(), widget.NewSeparator(),
				container.NewVBox(
					container.NewHBox(widgets.Status, layout.NewSpacer()),
					container.NewPadded(),
					container.NewPadded(),
					container.NewGridWithColumns(1, sbox.Widget())),
			),
		),
	)

	myWindow.Resize(fyne.NewSize(700, 400))
	widgets.Progress.Hidden = true
	myWindow.ShowAndRun()

	tidyUp(config)
}

func tidyUp(config config) {
	config.SaveConfig()
	log.Println("Cleaned up")
}
