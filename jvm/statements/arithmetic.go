package statements

import (
	"fmt"
	"go-on-jvm/jvm/constantpool"
	jvmio "go-on-jvm/jvm/io"
	"go-on-jvm/jvm/opcodes"
)

type ArithmeticType int

const (
	AddInt ArithmeticType = iota
	AddLong
	AddFloat
	AddDouble

	SubInt
	SubLong
	SubFloat
	SubDouble

	MultiplyInt
	MultiplyLong
	MultiplyFloat
	MultiplyDouble
)

func (a ArithmeticType) OpCode() int {
	switch a {
	case AddInt:
		return opcodes.IADD
	case AddLong:
		return opcodes.LADD
	case AddFloat:
		return opcodes.FADD
	case AddDouble:
		return opcodes.DADD
	case SubInt:
		return opcodes.ISUB
	case SubLong:
		return opcodes.LSUB
	case SubFloat:
		return opcodes.FSUB
	case SubDouble:
		return opcodes.DSUB
	case MultiplyInt:
		return opcodes.IMUL
	case MultiplyLong:
		return opcodes.LMUL
	case MultiplyFloat:
		return opcodes.FMUL
	case MultiplyDouble:
		return opcodes.DMUL
	default:
		panic(fmt.Errorf("unknown arithmetic type: %d", a))
	}
}

func (a ArithmeticType) New(left Statement, right Statement) Arithmetic {
	return Arithmetic{Type: a, Left: left, Right: right}
}

type Arithmetic struct {
	Type  ArithmeticType
	Left  Statement
	Right Statement
}

func (a Arithmetic) GetInstructions(stack *Stack, pool *constantpool.ConstantPool) []byte {
	instructions := a.Left.GetInstructions(stack, pool)
	instructions = append(instructions, a.Right.GetInstructions(stack, pool)...)
	instructions = jvmio.AppendPaddedBytes(instructions, a.Type.OpCode(), 1)
	return instructions
}

func (a Arithmetic) FillConstantsPool(pool *constantpool.ConstantPool) {
	a.Left.FillConstantsPool(pool)
	a.Right.FillConstantsPool(pool)
}

func (a Arithmetic) MaxStack() int {
	return a.Left.MaxStack() + a.Right.MaxStack()
}
