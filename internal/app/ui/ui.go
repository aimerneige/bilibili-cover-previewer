package ui

import (
	"github.com/aimerneige/bilibili-cover-previewer/internal/app/ui/bind"
	"github.com/webview/webview"
)

const (
	title  = "Bilibili Cover Previewer"
	width  = 400
	height = 720
)

// ShowMainWindow show main window
func ShowMainWindow(debug bool, port string) {
	w := webview.New(debug)
	defer w.Destroy()

	w.SetTitle(title)
	w.SetSize(width, height, webview.HintNone)
	bind.AllBindCollection(w)
	w.Navigate("http://localhost" + port)
	w.Run()
}
