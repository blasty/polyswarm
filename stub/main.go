// Invokes the perigord driver application

package main

import (
	_ "github.com/blasty/polyswarm/migrations"
	"github.com/blasty/perigord/stub"
)

func main() {
	stub.StubMain()
}
