package visitors

import (
	"fmt"
	"go-on-jvm/intermediate"
	"go/ast"
)

type variableGroupVisitor struct {
	VariableGroup intermediate.VariableGroup
	callback      visitedCallback
}

func (v *variableGroupVisitor) OnComplete(callback visitedCallback) {
	v.callback = callback
}

func (v *variableGroupVisitor) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return runCallback(v, v.callback)
	}

	switch node := node.(type) {
	case *ast.ValueSpec:
		visitor := variableDeclarationVisitor{}
		visitor.OnComplete(func(visitor astVisitor) ast.Visitor {
			v.VariableGroup.AddVariable(visitor.(*variableDeclarationVisitor).Variable)
			return v
		})
		return &visitor

	default:
		panic(fmt.Errorf("unexpected node given to variable group visitor: %#v", node))
	}

	return v
}

type variableDeclarationVisitor struct {
	Variable intermediate.Variable
	callback visitedCallback
}

func (v *variableDeclarationVisitor) OnComplete(callback visitedCallback) {
	v.callback = callback
}

func (v *variableDeclarationVisitor) Visit(node ast.Node) (w ast.Visitor) {
	if _, ok := node.(*ast.ValueSpec); node == nil || ok {
		// TODO: Handle in switch statement?
		return runCallback(v, v.callback)
	}

	switch node := node.(type) {
	case *ast.Ident:
		v.Variable.Name = node.Name

	case *ast.BasicLit:
		v.Variable.Value = node.Value

	default:
		panic(fmt.Errorf("unexpected node given to variable visitor: %#v", node))
	}

	return v
}
