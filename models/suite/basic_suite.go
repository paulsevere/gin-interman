package suite

import (
	"fmt"

	"github.com/paulsevere/gin-interman/interceptor/matcher"
	"github.com/paulsevere/gin-interman/models/stub"
)

// BasicSuite asdfs
type BasicSuite struct {
	Config    map[string]interface{} `json:"config"`
	Stubs     map[string]stub.Stub   `json:"stubs"`
	SuiteName string                 `json:"name"`
}

// NewBasicSuite asdfs
func NewBasicSuite(name string) *BasicSuite {

	return &BasicSuite{Config: make(map[string]interface{}), Stubs: make(map[string]stub.Stub), SuiteName: name}
}

func createMatchString(m matcher.Matcher) string {
	matchString := fmt.Sprintf("%s##%s", m["method"], m["path"])
	fmt.Println(matchString)
	return matchString

}

// MatchStub asdf
func (suite *BasicSuite) MatchStub(m matcher.Matcher) stub.Stub {
	return suite.Stubs[createMatchString(m)]
}

// StubList asdfasd
func (suite *BasicSuite) StubList() []string {
	keys := make([]string, len(suite.Stubs))
	i := 0
	for k := range suite.Stubs {
		keys[i] = k
		i++
	}
	return keys
}

// AddStub asdf
func (suite *BasicSuite) AddStub(s stub.Stub) error {
	suite.Stubs[s.Name()] = s
	return nil
}

// Name asdfasdf
func (suite *BasicSuite) Name() (name string) {
	name = suite.SuiteName
	return
}
