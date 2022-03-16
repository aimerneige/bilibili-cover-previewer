package web

import "github.com/gin-gonic/gin"

// StartWebServer start web server
func StartWebServer(port, root string) {
	r := gin.Default()

	appDir := root + "app/"
	staticDir := root + "static/"

	r.LoadHTMLFiles(appDir + "index.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	r.Static("/static", staticDir)

	r.Run(port)
}
