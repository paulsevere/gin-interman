package api

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/paulsevere/gin-interman/models/manager"
)

// SuitesRouter asdfs
func SuitesRouter(app *gin.RouterGroup) {

	app.POST("/suites", func(c *gin.Context) {
		m := c.MustGet("manager")
		man := m.(*manager.Manager)
		spew.Dump(man)
		var s interface{}

		if err := c.ShouldBindJSON(&s); err == nil {

			spew.Dump(s)
			man.Dispatch(manager.Action{Kind: "ADD_SUITE", Payload: s, Suite: ""})
			c.JSON(200, s)

		} else {

			c.String(404, err.Error())
			c.Abort()
			return
		}

	})
}
