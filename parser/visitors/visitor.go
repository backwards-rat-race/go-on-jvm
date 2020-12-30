package visitors

import (
	"fmt"
	"go-on-jvm/parser/structure"
	"go/ast"
)

type visitedCallback func(visitor astVisitor) ast.Visitor

type astVisitor interface {
	ast.Visitor
	OnComplete(visitedCallback)
}

func runCallback(visitor astVisitor, callback visitedCallback) ast.Visitor {
	if callback == nil {
		return nil
	} else {
		return callback(visitor)
	}
}

type Visitor struct {
	Parsed structure.Parsed
}

func (v *Visitor) Visit(node ast.Node) (w ast.Visitor) {
	if node == nil {
		return v
	}

	_, ok := node.(*ast.Package)
	if !ok {
		panic(fmt.Errorf("unexpected node given to root visitor: %#v", node))
	}

	visitor := packageVisitor{}
	visitor.OnComplete(func(visitor astVisitor) ast.Visitor {
		v.Parsed.AddPackage(visitor.(*packageVisitor).Package)
		return v
	})
	return &visitor
}

func New() *Visitor {
	return &Visitor{}
}
