package visitors

import (
	"fmt"
	"go-on-jvm/parser/structure"
	"go/ast"
)

type packageVisitor struct {
	Package  structure.Package
	callback visitedCallback
}

func (p *packageVisitor) OnComplete(callback visitedCallback) {
	p.callback = callback
}

func (p *packageVisitor) Visit(node ast.Node) ast.Visitor {
	// Completed parsing of this package
	if node == nil {
		return runCallback(p, p.callback)
	}

	_, ok := node.(*ast.File)
	if !ok {
		panic(fmt.Errorf("unexpected node given to package visitor: %#v", node))
	}

	visitor := fileVisitor{}
	visitor.OnComplete(func(visitor astVisitor) ast.Visitor {
		p.Package.AddDeclarationsContext(visitor.(*fileVisitor).DeclarationContext)
		return p
	})

	return &visitor
}
