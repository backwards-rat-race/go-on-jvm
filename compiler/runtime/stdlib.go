package runtime

import (
	"go-on-jvm/jvm"
	"go-on-jvm/jvm/definitions"
	"go-on-jvm/jvm/statements"
)

const StandardLibraryClassName = "StandardLibrary"
const AppendMethod = "append"

func NewStandardLib() definitions.Class {
	class := definitions.NewClass(StandardLibraryClassName, jvm.ObjectClass)
	class.AddMethod(constructor())
	class.AddMethod(append())
	return class
}

func constructor() definitions.Method {
	constructor := definitions.NewMethod(jvm.ConstructorName, definitions.Private)
	constructor.WithTypeDescriptor(jvm.Void)

	superCall := statements.NewInvocation(
		statements.NewMethodReference(jvm.ObjectClass, jvm.ConstructorName, jvm.Void),
		statements.SelfReferenceVariable,
	)
	constructor.AddStatement(superCall)
	constructor.AddStatement(statements.ReturnVoid)

	return constructor
}

func append() definitions.Method {
	append := definitions.NewMethod(AppendMethod, definitions.Public, definitions.Static)

	return append
}
