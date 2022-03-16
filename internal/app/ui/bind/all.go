package bind

import "github.com/webview/webview"

// AllBindCollection is a collection of all bind functions
func AllBindCollection(w webview.WebView) {
	w.Bind("preview", OpenBilibiliHomePageBind())
}
