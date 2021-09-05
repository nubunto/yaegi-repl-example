package main

import (
	"fmt"
	"go/build"
	"log"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

var Symbols interp.Exports = make(interp.Exports)

func main() {
	pkgs := setupSymbols()

	i := interp.New(interp.Options{GoPath: build.Default.GOPATH})

	// Import the standard library.
	if err := i.Use(stdlib.Symbols); err != nil {
		log.Fatal(err)
	}

	if err := i.Use(Symbols); err != nil {
		log.Fatal(err)
	}

	for _, pkg := range pkgs {
		i.Eval(fmt.Sprintf(`import "%s"`, pkg))
	}

	_, err := i.REPL()
	if err != nil {
		log.Fatal(err)
	}
}
