package stub

import (
	"encoding/json"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
)

// BodyStub simplest
type BodyStub struct {
	Suite    string `json:"suite"`
	Path     string `json:"path"`
	Method   string `json:"method"`
	Response string `json:"response"`
}

// NewBodyStub asdfas
func NewBodyStub(suite, path, method, response string) *BodyStub {
	return &BodyStub{Suite: suite, Path: path, Method: method, Response: response}
}

// Match sds
// func (stub *BodyStub) Match(matcher.Matcher) bool {
// 	return false
// 	// return stub.Method == m && stub.Path == p
// }

// Respond sd
func (stub *BodyStub) Respond(c *gin.Context) {
	spew.Dump("OK COOL")
	c.String(200, "%s", stub.Response)
	c.Abort()
}

// Serialize asdfa
func (stub *BodyStub) Serialize() []byte {
	jsonStub, err := json.Marshal(stub)
	if err != nil {
		return []byte("ERROR")
	}
	return jsonStub
}

// Name asdf
func (stub *BodyStub) Name() string {
	return fmt.Sprintf("%s##%s", stub.Method, stub.Path)

}
