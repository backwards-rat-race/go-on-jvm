package visitors

import (
	"fmt"
	"go-on-jvm/parser/structure"
	"go/ast"
)

type importDeclarationVisitor struct {
	Import   structure.Import
	callback visitedCallback
}

func (i *importDeclarationVisitor) OnComplete(callback visitedCallback) {
	i.callback = callback
}

func (i *importDeclarationVisitor) Visit(node ast.Node) (w ast.Visitor) {
	if node == nil {
		return runCallback(i, i.callback)
	}

	switch node := node.(type) {
	case *ast.ImportSpec:
		// Ignored

	case *ast.Ident:
		i.Import.Alias = node.Name

	case *ast.BasicLit:
		i.Import.Package = node.Value

	default:
		panic(fmt.Errorf("unexpected node given to import visitor: %#v", node))
	}

	return i
}
