package stub

import (
	"github.com/gin-gonic/gin"
)

// Stub -- stub interace
type Stub interface {
	// Match(matcher.Matcher) bool
	Respond(c *gin.Context)
	Serialize() []byte
	Name() string
}
