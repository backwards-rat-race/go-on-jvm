package runtime

import (
	"go-on-jvm/jvm"
)

const StandardLibraryClassName = "StandardLibrary"

func NewStandardLib() jvm.Class {
	class := jvm.NewClass(StandardLibraryClassName, jvm.ObjectClass)
	class.AddMethod(constructor())
	return class
}

func constructor() jvm.Method {
	constructor := jvm.NewMethod(jvm.ConstructorName, jvm.Private)
	constructor.WithTypeDescriptor(jvm.Void)

	superCall := jvm.NewInvocationStatement(
		jvm.NewMethodReference(jvm.ObjectClass, jvm.ConstructorName, jvm.Void),
		jvm.SelfReferenceVariable,
	)
	constructor.AddStatement(superCall)
	constructor.AddStatement(jvm.NewReturnStatement())

	return constructor
}
