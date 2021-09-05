package main

import (
	"reflect"

	"github.com/nubunto/repl-example/bar"
	"github.com/nubunto/repl-example/foo"
)

func setupSymbols() []string {
	fooCap := foo.NewCapability()
	barCap := bar.NewCapability()

	var pkgs []string

	pkgs = append(pkgs, "github.com/nubunto/repl-example/foo")
	Symbols["github.com/nubunto/repl-example/foo/foo"] = map[string]reflect.Value{
		"Capability": reflect.ValueOf(fooCap),
	}

	pkgs = append(pkgs, "github.com/nubunto/repl-example/bar")
	Symbols["github.com/nubunto/repl-example/bar/bar"] = map[string]reflect.Value{
		"Capability": reflect.ValueOf(barCap),
	}

	return pkgs
}
