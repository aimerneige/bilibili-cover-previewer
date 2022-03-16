package bind

import (
	"time"

	"github.com/webview/webview"
)

// inject js
var js string = `
document.getElementsByClassName("nav-search-input")[0].placeholder = "fuck";
`

// OpenBilibiliHomePageBind is a bind function for opening bilibili homepage
func OpenBilibiliHomePageBind() interface{} {
	return func() error {
		w := webview.NewWindow(false, nil)
		w.SetSize(400, 580, webview.HintNone)
		w.Navigate("https://www.bilibili.com/")
		w.Dispatch(func() {
			go func() {
				time.Sleep(time.Second * 4)
				w.Eval(js)
			}()
		})
		w.Run()
		return nil
	}
}
