package manager

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/paulsevere/gin-interman/models/stub"
	"github.com/paulsevere/gin-interman/models/suite"
)

// Action asdfs
type Action struct {
	Suite   string      `json:"suite"`
	Kind    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

// Manager assdfsfs
type Manager struct {
	Suites     map[string]suite.Suite
	SuiteList  []string
	Dispatches chan Action
}

// New asdfsf
func New() *Manager {
	sl := suite.NewBasicSuite("first")
	st := stub.NewBodyStub("first", "/home", "GET", "HELLO WORLD!")
	sl.AddStub(st)
	return &Manager{Suites: map[string]suite.Suite{"first": sl}, SuiteList: []string{"first"}, Dispatches: make(chan Action, 0)}
}

// Middleware asdfasfd
func Middleware(m *Manager) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Set("manager", m)
	}
}

// Dispatch asdfs
func (m *Manager) Dispatch(action Action) bool {
	spew.Dump(action)
	return true
}

// HandleQuery asdf
func (m *Manager) HandleQuery(name string) suite.Suite {
	fmt.Printf("IN QUERY HANDLER : %s\n", name)
	spew.Dump(m.Suites[name])
	if s := m.Suites[name]; s != nil {
		spew.Dump(s)
		return s
	}
	return nil
}

// HandleAction asdfasdfsf
func (m *Manager) HandleAction(a Action) bool {

	switch kind := a.Kind; kind {
	case "ADD_SUITE":
		if s, ok := a.Payload.(suite.Suite); ok {
			m.Suites[s.Name()] = s
			m.SuiteList = append()
		}
	case "REMOVE_SUITE":
		// delete()
	}
	return false
}

func (m *Manager) hasSuite(s string) bool {
	for _, su := range m.SuiteList {
		if su == s {
			return true
		}
	}
	return false
}
