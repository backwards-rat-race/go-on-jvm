package runtime

import (
	"go-on-jvm/jvm/definitions"
	"go-on-jvm/jvm/statements"
	"go-on-jvm/jvm/types"
)

const (
	StandardLibraryClassName = "StandardLibrary"
	AppendMethod             = "append"

	appendMethodOriginalArg   = "original"
	appendMethodAdditionalArg = "additional"
)

func NewStandardLib() definitions.Class {
	class := definitions.NewClass(StandardLibraryClassName, types.ObjectClass)
	class.AddMethod(constructorMethod())
	class.AddMethod(appendMethod())
	return class
}

func constructorMethod() definitions.Method {
	constructor := definitions.NewMethod(types.ConstructorName, definitions.Private)
	constructor.ReturnType = types.Void

	superCall := statements.NewInvocation(
		statements.NewMethodReference(types.ObjectClass, types.ConstructorName, types.Void),
		statements.GetSelfReferenceVariable,
	)
	constructor.AddStatement(superCall)
	constructor.AddStatement(statements.NewVoidReturn())

	return constructor
}

func appendMethod() definitions.Method {
	appendMethod := definitions.NewMethod(AppendMethod, definitions.Public, definitions.Static, definitions.VarArgs)
	originalArg := statements.NewLocalVariable(appendMethodOriginalArg, types.ObjectClass)
	additionalArg := statements.NewLocalVariable(appendMethodAdditionalArg, types.ObjectClass)
	appendMethod.AddArgument(originalArg)
	appendMethod.AddArgument(additionalArg)

	createdVar := statements.NewLocalVariable("created", types.ObjectClass.Array())

	originalIsNull := statements.NewBlock()
	originalIsNull.AddStatement(
		statements.NewVariableSet(
			createdVar,
			statements.NewStaticInvocation(
				statements.NewMethodReference(
					types.MustParse("java.util.Arrays"),
					"copyOf",
					types.ObjectClass.Array(),
					types.ObjectClass.Array(),
					types.Int,
				),
				statements.NewVariableGet(additionalArg),
				statements.NewArrayLen(
					statements.NewVariableGet(additionalArg),
				),
			),
		),
	)

	originalIsNotNull := statements.NewBlock()
	originalIsNotNull.AddStatement(
		statements.NewVariableSet(
			createdVar,
			statements.NewStaticInvocation(
				statements.NewMethodReference(
					types.MustParse("java.util.Arrays"),
					"copyOf",
					types.ObjectClass.Array(),
					types.ObjectClass.Array(),
					types.Int,
				),
				statements.NewVariableGet(originalArg),
				statements.AddInt.New(
					statements.NewArrayLen(
						statements.NewVariableGet(originalArg),
					),
					statements.NewArrayLen(
						statements.NewVariableGet(additionalArg),
					),
				),
			),
		),
	)
	originalIsNotNull.AddStatement(
		statements.NewStaticInvocation(
			statements.NewMethodReference(
				types.MustParse("System"),
				"arraycopy",
				types.Void,
				types.ObjectClass,
				types.Int,
				types.ObjectClass,
				types.Int,
				types.Int,
			),
			statements.NewVariableGet(additionalArg),
			statements.NewIntConstant(0),
			statements.NewVariableGet(createdVar),
			statements.NewArrayLen(
				statements.NewVariableGet(originalArg),
			),
			statements.NewArrayLen(
				statements.NewVariableGet(additionalArg),
			),
		),
	)

	ifOriginalIsNull := statements.IfNull.New(statements.NewVariableGet(originalArg))
	ifOriginalIsNull.Success = originalIsNull
	ifOriginalIsNull.Failure = originalIsNotNull
	appendMethod.AddStatement(ifOriginalIsNull)

	appendMethod.AddStatement(
		statements.NewReturn(
			statements.ReturnReference,
			statements.NewVariableGet(createdVar),
		),
	)

	return appendMethod
}
