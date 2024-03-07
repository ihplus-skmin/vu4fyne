package main

import (
	//"image/color"
	//"time"
	//"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	//"fyne.io/fyne/v2/canvas"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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
	var data = []string{"a", "string", "listwdwwdwdwdwfwfwwfwfwfwfwfwwfw"}

	myApp := app.New()
	myWindow := myApp.NewWindow("Video Clip Uploader")

	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		})
	btn1 := widget.NewButton("click me", uploading)
	btn2 := widget.NewButton("btn2", func() {})
	btn3 := widget.NewButton("btn3", func() {})

	content := container.New(layout.NewHBoxLayout(), list)
	grid := container.New(layout.NewHBoxLayout(), btn1, btn2, btn3)

	myWindow.SetContent(container.New(layout.NewVBoxLayout(), content, grid))
	myWindow.Resize(fyne.NewSize(500, 200))
	list.Resize(fyne.NewSize(500, 200))
	myWindow.ShowAndRun()

	tidyUp()
}

func tidyUp() {

}
