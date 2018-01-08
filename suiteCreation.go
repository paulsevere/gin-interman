package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/paulsevere/gin-interman/models/stub"
	"github.com/paulsevere/gin-interman/models/suite"
)

func tester() string {
	sl := suite.NewBasicSuite("first")
	st := stub.NewBodyStub("first", "/home", "GET", "HELLO WORLD!")
	sl.AddStub(st)
	stub := sl.MatchStub(map[string]string{"path": "/home", "method": "GET"})

	spew.Dump(stub)
	return "ok"
}
