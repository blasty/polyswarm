package main

import (
	"testing"

	_ "github.com/blasty/polyswarm/tests"
	_ "github.com/blasty/polyswarm/migrations"
	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner
func Test(t *testing.T) { TestingT(t) }
