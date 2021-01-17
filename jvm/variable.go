package jvm

type Variable struct {
	Object string
	Name   string
	Type   string
}

func NewLocalVariable(name string, typeDescriptor string) Variable {
	return Variable{Name: name, Type: typeDescriptor}
}

func NewObjectVariable(object string, name string, typeDescriptor string) Variable {
	return Variable{object, name, typeDescriptor}
}

func (v Variable) IsLocal() bool {
	return v.Object == ""
}

var SelfReferenceVariable = Variable{}
