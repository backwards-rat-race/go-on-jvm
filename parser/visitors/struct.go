package visitors

import (
	"fmt"
	"go-on-jvm/intermediate"
	"go/ast"
)

type structVisitor struct {
	Class    intermediate.Class
	callback visitedCallback
}

func (s *structVisitor) OnComplete(callback visitedCallback) {
	s.callback = callback
}

func (s *structVisitor) Visit(node ast.Node) (w ast.Visitor) {
	if node == nil {
		return runCallback(s, s.callback)
	}

	switch node.(type) {
	case *ast.FieldList:
		visitor := fieldListVisitor{}
		visitor.OnComplete(func(visitor astVisitor) ast.Visitor {
			s.Class.AddFields(visitor.(*fieldListVisitor).Fields)
			return s
		})
		return &visitor

	default:
		panic(fmt.Errorf("unexpected node given to struct visitor: %#v", node))
	}

	return s
}
