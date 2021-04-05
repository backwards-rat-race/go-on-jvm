package runtime

import (
	"go-on-jvm/jvm/definitions"
	"go-on-jvm/jvm/statements"
	"go-on-jvm/jvm/types"
)

const (
	StandardLibraryClassName = "StandardLibrary"
	AppendMethod             = "append"
	JoinMethod               = "join"
	PrintlnMethod            = "println"
)

func NewStandardLib() definitions.Class {
	class := definitions.NewClass(StandardLibraryClassName, types.ObjectClass)
	class.AddMethod(constructorMethod())
	class.AddMethod(printlnMethod())
	class.AddMethod(appendMethod())
	class.AddMethod(joinMethod())
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

func printlnMethod() definitions.Method {
	printlnMethod := definitions.NewMethod(PrintlnMethod, definitions.Public, definitions.Static, definitions.VarArgs)
	printlnMethod.ReturnType = types.Void
	objectsArg := statements.NewLocalVariable("objects", types.ObjectClass.Array())
	printlnMethod.AddArgument(objectsArg)

	stringClass := types.MustParse("java.lang.String")

	printlnReference := statements.NewMethodReference(
		types.MustParse("java.io.PrintStream"),
		"println",
		types.Void,
		stringClass,
	)
	joinReference := statements.NewMethodReference(
		types.MustParse(StandardLibraryClassName),
		"join",
		stringClass,
		types.ObjectClass.Array(),
	)
	printlnMethod.AddStatement(
		statements.NewInvocation(
			printlnReference,
			statements.NewStaticVariableGet(
				statements.NewStaticVariable(
					"java/lang/System",
					"out",
					types.MustParse("java.io.PrintStream"),
				),
			),
			statements.NewStaticInvocation(
				joinReference,
				statements.NewVariableGet(objectsArg),
			),
		),
	)
	printlnMethod.AddStatement(statements.NewVoidReturn())
	return printlnMethod
}

func appendMethod() definitions.Method {
	appendMethod := definitions.NewMethod(AppendMethod, definitions.Public, definitions.Static, definitions.VarArgs)
	appendMethod.ReturnType = types.ObjectClass.Array()
	originalArg := statements.NewLocalVariable("original", types.ObjectClass)
	additionalArg := statements.NewLocalVariable("additional", types.ObjectClass.Array())
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

	ifOriginalIsNull := statements.NewIf(
		statements.IsNull.New(
			statements.NewVariableGet(originalArg),
		),
	)
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

func joinMethod() definitions.Method {
	stringClass := types.MustParse("java.lang.String")

	joinMethod := definitions.NewMethod(JoinMethod, definitions.Private, definitions.Static, definitions.VarArgs)
	joinMethod.ReturnType = stringClass
	objectsArg := statements.NewLocalVariable("objects", types.ObjectClass.Array())
	joinMethod.AddArgument(objectsArg)

	stringJoinerClass := types.MustParse("java.util.StringJoiner")
	stringJoiner := statements.NewLocalVariable("sj", stringJoinerClass)
	joinMethod.AddStatement(
		statements.NewVariableSet(
			stringJoiner,
			statements.NewInitInvocation(
				stringJoinerClass,
				statements.NewStringConstant(" "),
			),
		),
	)

	forIndex := statements.NewLocalVariable("i", types.Int)
	forBlock := statements.NewBlock()

	ifEntryIsNull := statements.NewIf(
		statements.IsNull.New(
			statements.NewArrayGet(
				statements.NewVariableGet(objectsArg),
				statements.NewVariableGet(forIndex),
			),
		),
	)
	ifEntryIsNull.Success.AddStatement(
		statements.NewThrowAway(
			statements.NewInvocation(
				statements.NewMethodReference(
					stringJoinerClass,
					"add",
					stringJoinerClass,
					stringClass,
				),
				statements.NewVariableGet(
					stringJoiner,
				),
				statements.NewStringConstant("0x00"),
			),
		),
	)
	ifEntryIsNull.Failure.AddStatement(
		statements.NewThrowAway(
			statements.NewInvocation(
				statements.NewMethodReference(
					stringJoinerClass,
					"add",
					stringJoinerClass,
					stringClass,
				),
				statements.NewVariableGet(
					stringJoiner,
				),
				statements.NewInvocation(
					statements.NewMethodReference(
						types.ObjectClass.Array(),
						"toString",
						stringJoinerClass,
					),
					statements.NewArrayGet(
						statements.NewVariableGet(objectsArg),
						statements.NewVariableGet(objectsArg),
					),
				),
			),
		),
	)

	forBlock.AddStatement(ifEntryIsNull)

	joinMethod.AddStatement(
		statements.NewForLoop(
			forBlock,
			statements.NewVariableSet(
				forIndex,
				statements.NewIntConstant(0),
			),
			statements.IsLessThan.New(
				statements.SubInt.New(
					statements.NewVariableGet(forIndex),
					statements.NewArrayLen(
						statements.NewVariableGet(objectsArg),
					),
				),
			),
			statements.NewVariableSet(
				forIndex,
				statements.AddInt.New(
					statements.NewVariableGet(forIndex),
					statements.NewIntConstant(1),
				),
			),
		),
	)
	joinMethod.AddStatement(
		statements.NewReturn(
			statements.ReturnReference,
			statements.NewInvocation(
				statements.NewMethodReference(
					stringJoinerClass,
					"toString",
					stringClass,
				),
				statements.NewVariableGet(
					stringJoiner,
				),
			),
		),
	)

	return joinMethod
}
