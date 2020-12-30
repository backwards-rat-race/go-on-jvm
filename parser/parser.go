package parser

import (
	"go-on-jvm/parser/structure"
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

func ParseDirectory(path string) (structure.Parsed, error) {
	fileSet := token.NewFileSet()
	pkgs, err := parser.ParseDir(fileSet, path, nil, parser.AllErrors)

	if err != nil {
		return structure.Parsed{}, err
	}

	visitor := visitors.New()

	for _, a := range pkgs {
		ast.Walk(visitor, a)
	}

	return visitor.Parsed, nil
}
