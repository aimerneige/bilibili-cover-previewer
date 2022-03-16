package ui

import (
	"github.com/webview/webview"
)

const (
	title  = "Bilibili Cover Previewer"
	width  = 1024
	height = 768
)

// ShowMainWindow show main window
func ShowMainWindow(debug bool, port string) {
	w := webview.New(debug)
	defer w.Destroy()

	w.SetTitle(title)
	w.SetSize(width, height, webview.HintNone)

	w.Navigate("http://localhost" + port)
	w.Run()
}
