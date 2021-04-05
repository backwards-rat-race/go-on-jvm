package statements

import (
	"go-on-jvm/jvm/constantpool"
	jvmio "go-on-jvm/jvm/io"
	"go-on-jvm/jvm/opcodes"
	jvmtypes "go-on-jvm/jvm/types"
)

type MethodReference struct {
	Class jvmtypes.TypeReference
	Name  string
	Type  jvmtypes.MethodType
}

type MethodType struct {
	Arguments []string
	Return    string
}

func NewMethodReference(class jvmtypes.TypeReference, name string, returnType jvmtypes.TypeReference, args ...jvmtypes.TypeReference) MethodReference {
	return MethodReference{
		Class: class,
		Name:  name,
		Type:  jvmtypes.MethodType{ReturnType: returnType, Arguments: args},
	}
}

func (m MethodReference) IsVoid() bool {
	return m.Type.ReturnType == jvmtypes.Void
}

type Invocation struct {
	MethodReference MethodReference
	Static          bool
	Vars            []Statement
}

func (i Invocation) GetInstructions(stack *Stack, pool *constantpool.ConstantPool) []byte {
	instructions := make([]byte, 0)

	for _, variable := range i.Vars {
		instructions = append(instructions, variable.GetInstructions(stack, pool)...)
	}

	instructions = jvmio.AppendPaddedBytes(instructions, i.opcode(), 1)

	constantIndex := pool.FindMethodReference(i.MethodReference.Class.Jvm(), i.MethodReference.Name, i.MethodReference.Type.Descriptor())
	return jvmio.AppendPaddedBytes(instructions, constantIndex, 2)
}

func (i Invocation) FillConstantsPool(pool *constantpool.ConstantPool) {
	pool.AddMethodReference(i.MethodReference.Class.Jvm(), i.MethodReference.Name, i.MethodReference.Type.Descriptor())
	for _, statement := range i.Vars {
		statement.FillConstantsPool(pool)
	}
}

func (i Invocation) MaxStack() uint {
	max := uint(len(i.Vars))

	for _, statement := range i.Vars {
		max = iMax(max, statement.MaxStack())
	}

	return max
}

func (i Invocation) opcode() uint {
	if i.Static {
		return opcodes.INVOKESTATIC
	} else {
		return opcodes.INVOKESPECIAL
	}
}

func NewInvocation(method MethodReference, vars ...Statement) Invocation {
	return Invocation{
		MethodReference: method,
		Vars:            vars,
	}
}

func NewStaticInvocation(method MethodReference, vars ...Statement) Invocation {
	return Invocation{
		MethodReference: method,
		Static:          true,
		Vars:            vars,
	}
}
