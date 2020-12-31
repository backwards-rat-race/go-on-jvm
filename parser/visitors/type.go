package visitors

import (
	"fmt"
	"go-on-jvm/intermediate"
	"go/ast"
)

type visitedStructCallback func(class intermediate.Class)

type typeVisitor struct {
	Name           string
	callback       visitedCallback
	structCallback visitedStructCallback
}

func (t *typeVisitor) OnComplete(callback visitedCallback) {
	t.callback = callback
}

func (t *typeVisitor) OnCompleteStruct(callback visitedStructCallback) {
	t.structCallback = callback
}

func (t *typeVisitor) Visit(node ast.Node) (w ast.Visitor) {
	if node == nil {
		return runCallback(t, t.callback)
	}

	switch node := node.(type) {
	case *ast.TypeSpec:
		//fmt.Printf("Type: TypeSpec Node '%s' %#v\n", node.Name, node)

	case *ast.Ident:
		t.Name = node.Name

	case *ast.StructType:
		fmt.Printf("Type: Struct Node %#v\n", node)
		visitor := structVisitor{}
		visitor.Class.Name = t.Name
		visitor.OnComplete(func(visitor astVisitor) ast.Visitor {
			if t.structCallback != nil {
				t.structCallback(visitor.(*structVisitor).Class)
			}
			return t
		})
		return &visitor

	default:
		panic(fmt.Errorf("unexpected node given to type visitor: %#v", node))
	}

	return t
}
