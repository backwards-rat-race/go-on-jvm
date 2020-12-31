package visitors

import (
	"fmt"
	"go-on-jvm/intermediate"
	"go/ast"
	"go/token"
)

type fileVisitor struct {
	Encapsulated intermediate.Encapsulated
	callback     visitedCallback
}

func (f *fileVisitor) OnComplete(callback visitedCallback) {
	f.callback = callback
}

func (f *fileVisitor) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return runCallback(f, f.callback)
	}

	switch node := node.(type) {
	case *ast.File:
		// Skip
		//fmt.Printf("File: File Node %#v\n", node)

	case *ast.GenDecl:
		return newGenericDeclarationVisitor(f, node, &f.Encapsulated)

	case *ast.Ident:
		f.Encapsulated.Package = node.Name
		fmt.Printf("File: Ident Node %#v\n", node)
		return nil

	case *ast.FuncDecl:
		fmt.Printf("File: FuncDecl Node %#v\n", node)
		visitor := funcVisitor{}
		visitor.callback = func(visitor astVisitor) ast.Visitor {
			return f
		}
		return &visitor

	case *ast.FuncType:
		fmt.Printf("File: FuncType Node %#v\n", node)
		return nil

	default:
		panic(fmt.Errorf("unexpected node given to file visitor: %#v", node))
	}

	return f
}

func newGenericDeclarationVisitor(parent ast.Visitor, node *ast.GenDecl, context *intermediate.Encapsulated) ast.Visitor {
	switch node.Tok {
	case token.IMPORT:
		visitor := importDeclarationVisitor{}
		visitor.OnComplete(func(visitor astVisitor) ast.Visitor {
			context.AddImport(visitor.(*importDeclarationVisitor).Import)
			return parent
		})
		return &visitor

	case token.CONST:
		fallthrough
	case token.VAR:
		visitor := variableGroupVisitor{}
		visitor.VariableGroup.Const = node.Tok == token.CONST
		visitor.OnComplete(func(visitor astVisitor) ast.Visitor {
			context.AddVariableGroup(visitor.(*variableGroupVisitor).VariableGroup)
			return parent
		})
		return &visitor

	case token.TYPE:
		visitor := typeVisitor{}
		visitor.OnCompleteStruct(func(class intermediate.Class) {
			context.AddClass(class)
		})
		visitor.OnComplete(func(visitor astVisitor) ast.Visitor {
			return parent
		})
		return &visitor
	default:
		fmt.Printf("Declaration: Unknown Node %#v\n", node)
	}

	return nil
}
