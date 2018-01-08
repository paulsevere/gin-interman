package suite

import (
	"github.com/paulsevere/gin-interman/interceptor/matcher"
	"github.com/paulsevere/gin-interman/models/stub"
)

// Suite ok
type Suite interface {
	MatchStub(matcher.Matcher) stub.Stub
	StubList() []string
	AddStub(s stub.Stub) error
	Name() string
}

// List asdfs
type List []Suite
