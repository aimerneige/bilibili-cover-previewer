package main

import (
	"github.com/aimerneige/bilibili-cover-previewer/internal/app/ui"
	"github.com/aimerneige/bilibili-cover-previewer/internal/app/web"
)

// port and root can be overrides during build
// See the Makefile for more information.
var port = ":8080"
var root = "./public"

func main() {
	go web.StartWebServer(port, root)
	ui.ShowMainWindow(true, port)
}
