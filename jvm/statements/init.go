package statements

import (
	"go-on-jvm/jvm/constantpool"
	jvmio "go-on-jvm/jvm/io"
	"go-on-jvm/jvm/opcodes"
	jvmtypes "go-on-jvm/jvm/types"
)

type InitInvocation struct {
	Class jvmtypes.TypeReference
	Args  []Statement
}

func (i InitInvocation) GetInstructions(stack *Stack, pool *constantpool.ConstantPool) []byte {
	instructions := []byte{opcodes.NEW}
	index := pool.FindClassNameItem(i.Class.Jvm())
	instructions = jvmio.AppendPaddedBytes(instructions, index, 2)

	// We now have a reference to the new variable on stack. The first arg to init is also the variable,
	// so now we duplicate it
	instructions = append(instructions, opcodes.DUP)

	for _, variable := range i.Args {
		instructions = append(instructions, variable.GetInstructions(stack, pool)...)
	}

	initRef := NewMethodReference(i.Class, jvmtypes.ConstructorName, jvmtypes.Void)
	instructions = jvmio.AppendPaddedBytes(instructions, opcodes.INVOKESPECIAL, 1)
	constantIndex := pool.FindMethodReference(initRef.Class.Jvm(), initRef.Name, initRef.Type.Descriptor())
	instructions = jvmio.AppendPaddedBytes(instructions, constantIndex, 2)

	return instructions
}

func (i InitInvocation) FillConstantsPool(pool *constantpool.ConstantPool) {
	pool.AddClassReference(i.Class.Jvm())
	for _, statement := range i.Args {
		statement.FillConstantsPool(pool)
	}
}

func (i InitInvocation) MaxStack() uint {
	max := uint(len(i.Args))

	for _, statement := range i.Args {
		max = iMax(max, statement.MaxStack())
	}

	return max
}

func NewInitInvocation(class jvmtypes.TypeReference, args ...Statement) InitInvocation {
	return InitInvocation{Class: class, Args: args}
}
