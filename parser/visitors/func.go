package visitors

import (
	"fmt"
	"go/ast"
)

type funcVisitor struct {
	callback visitedCallback
}

func (f *funcVisitor) Visit(node ast.Node) (w ast.Visitor) {
	if node == nil {
		return runCallback(f, f.callback)
	}

	fmt.Printf("Func: node %#v\n", node)

	return f
}

func (f *funcVisitor) OnComplete(callback visitedCallback) {
	f.callback = callback
}
