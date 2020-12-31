package parser

import (
	"go-on-jvm/intermediate"
	"go-on-jvm/parser/visitors"
	"go/ast"
	"go/parser"
	"go/token"
)

type Parser struct {
	fileSet *token.FileSet
}

func New() *Parser {
	return &Parser{token.NewFileSet()}
}

func ParseDirectory(path string) (p intermediate.Parsed, err error) {
	fileSet := token.NewFileSet()
	pkgs, err := parser.ParseDir(fileSet, path, nil, parser.AllErrors)

	if err != nil {
		return
	}

	defer func() {
		r := recover()

		if r == nil {
			return
		}

		recoveredErr, ok := r.(error)

		if ok {
			err = recoveredErr
		} else {
			panic(r)
		}
	}()

	visitor := visitors.New()
	for _, a := range pkgs {
		ast.Walk(visitor, a)
	}
	p = visitor.Parsed

	return
}
