package compiler

import (
	"go-on-jvm/intermediate"
	"go-on-jvm/jvm"
)

type CompiledClass struct {
	Path  string
	Class jvm.Class
}

func Compile(_ intermediate.Parsed) []CompiledClass {
	return nil
}
