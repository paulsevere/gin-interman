package interceptor

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/paulsevere/gin-interman/interceptor/matcher"
	"github.com/paulsevere/gin-interman/models/manager"
)

// Initialize asdfas
func Initialize(m *manager.Manager) func(c *gin.Context) {
	return func(c *gin.Context) {
		Interceptor(c, m)
	}
}

// MatchMaker asdfs
func MatchMaker(c *gin.Context) {
	method := c.Request.Method
	path := c.Param("path")
	fmt.Printf("ROUTE:  %s , METHOD:  %s\n", path, method)
	c.Set("matcher", matcher.Matcher{"method": method, "path": path})
	spew.Dump(c.Get("matcher"))
	c.Next()
}

// Interceptor asdfas
func Interceptor(c *gin.Context, m *manager.Manager) {
	suiteName, err := c.Cookie("SUITE")
	fmt.Printf("SUITE NAME:  %s\n", suiteName)
	if err != nil {
		c.Next()
		return
	}
	suite := m.HandleQuery(suiteName)
	if suite == nil {
		c.Next()
		return
	}
	mat := c.GetStringMapString("matcher")
	if mat == nil {
		c.Next()
	}
	if stub := suite.MatchStub(mat); stub != nil {
		stub.Respond(c)
		c.Abort()
		return
	}

	// stub1.Respond(c)

	c.Next()
}
