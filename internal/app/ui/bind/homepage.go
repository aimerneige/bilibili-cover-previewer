package bind

import (
	"fmt"
	"time"

	"github.com/webview/webview"
)

// inject js
var js string = `
document.getElementsByClassName("bili-video-card__info--tit")[0].innerHTML = "%s";
document.getElementsByClassName("bili-video-card__info--author")[0].innerHTML = "%s";
document.getElementsByClassName("bili-video-card__info--date")[0].innerHTML = "· %s";
document.getElementsByClassName("bili-video-card__stats--text")[0].innerHTML = "%s";
document.getElementsByClassName("bili-video-card__stats--text")[1].innerHTML = "%s";
document.getElementsByClassName("bili-video-card__stats__duration")[0].innerHTML = "%s";
document.getElementsByClassName("bili-video-card__cover")[0].getElementsByTagName("source")[0].remove();
document.getElementsByClassName("bili-video-card__cover")[0].getElementsByTagName("img")[0].removeAttribute("loading");
document.getElementsByClassName("bili-video-card__cover")[0].getElementsByTagName("img")[0].removeAttribute("onload")
document.getElementsByClassName("bili-video-card__cover")[0].getElementsByTagName("img")[0].src = "%s";
document.getElementsByClassName("bili-video-card__cover")[0].getElementsByTagName("img")[0].alt = "%s";
`

// OpenBilibiliHomePageBind is a bind function for opening bilibili homepage
func OpenBilibiliHomePageBind() interface{} {
	return func(title, author, date, play, star, duration, cover, alt string) error {
		w := webview.NewWindow(false, nil)
		w.SetSize(1366, 768, webview.HintNone)
		w.Navigate("https://www.bilibili.com/")
		w.Dispatch(func() {
			go func() {
				time.Sleep(time.Second * 4)
				w.Eval(fmt.Sprintf(js, title, author, date, play, star, duration, cover, alt))
			}()
		})

		w.Run()
		return nil
	}
}
