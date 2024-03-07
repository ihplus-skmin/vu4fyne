package main

import (
	//"image/color"
	//"time"
	//"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	//"fyne.io/fyne/v2/canvas"

	//"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

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
	myWindow := myApp.NewWindow("Video Clip Uploader")

	content := widget.NewButton("click me", uploading)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(500, 200))
	myWindow.ShowAndRun()

	tidyUp()
}

func tidyUp() {

}
