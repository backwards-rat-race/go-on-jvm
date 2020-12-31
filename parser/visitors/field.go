package visitors

import (
	"fmt"
	"go-on-jvm/intermediate"
	"go/ast"
)

type fieldListVisitor struct {
	Fields   []intermediate.Field
	callback visitedCallback
}

func (f *fieldListVisitor) OnComplete(callback visitedCallback) {
	f.callback = callback
}

func (f *fieldListVisitor) Visit(node ast.Node) (w ast.Visitor) {
	if node == nil {
		return runCallback(f, f.callback)
	}

	switch node := node.(type) {
	case *ast.Field:
		f.addField(node)
		return nil

	default:
		panic(fmt.Errorf("unexpected node given to field list visitor: %#v", node))
	}

	return f
}

func (f *fieldListVisitor) addField(node *ast.Field) {
	field := intermediate.Field{}

	switch len(node.Names) {
	case 0:
		field.TypeOnly = true
	case 1:
		field.Name = node.Names[0].Name
	default:
		// TODO: multiple names?
		panic(fmt.Errorf("node has unexpected names count: %#v", node))
	}

	// TODO: support non basic types, e.g. anonymous struct
	typeNode, ok := node.Type.(*ast.Ident)
	if !ok {
		panic(fmt.Errorf("unexpected node given to field creation: %#v", node))
	}
	field.Type = typeNode.Name

	f.Fields = append(f.Fields, field)
}
