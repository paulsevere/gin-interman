package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	apiRouter "github.com/paulsevere/gin-interman/api"
	"github.com/paulsevere/gin-interman/cookie"
	"github.com/paulsevere/gin-interman/interceptor"
	"github.com/paulsevere/gin-interman/models/manager"
)

func main() {

	m := manager.New()
	app := gin.Default()
	app.Use(manager.Middleware(m))
	api := app.Group("/api")
	api.GET("set-cookie", cookie.SetSuiteCookie)
	apiRouter.SuitesRouter(api)
	proxy := app.Group("/proxy")
	proxy.Any("/*path", interceptor.MatchMaker, interceptor.Initialize(m), func(c *gin.Context) {
		c.String(200, "Coolio")
	})

	app.GET("/", func(c *gin.Context) {
		spew.Dump(c)
		c.String(200, "cool we are done")
	})

	app.Run(":3000")
}
