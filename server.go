package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/paulsevere/gin-interman/cookie"
)

func main() {
	app := gin.Default()

	api := app.Group("/api")
	api.GET("set-cookie", cookie.SetSuiteCookie)

	app.GET("/", func(c *gin.Context) {
		spew.Dump(c)
		c.String(200, "cool we are done")
	})

	app.Run(":3000")
}
