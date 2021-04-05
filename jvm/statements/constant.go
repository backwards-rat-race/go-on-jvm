package statements

import (
	"go-on-jvm/jvm/constantpool"
	jvmio "go-on-jvm/jvm/io"
	"go-on-jvm/jvm/opcodes"
)

type IntConstant struct {
	Constant int
}

func (i IntConstant) GetInstructions(stack *Stack, pool *constantpool.ConstantPool) []byte {
	var instructions []byte

	switch i.Constant {
	case -1:
		instructions = append(instructions, opcodes.ICONST_M1)
	case 0:
		instructions = append(instructions, opcodes.ICONST_0)
	case 1:
		instructions = append(instructions, opcodes.ICONST_1)
	case 2:
		instructions = append(instructions, opcodes.ICONST_2)
	case 3:
		instructions = append(instructions, opcodes.ICONST_3)
	case 4:
		instructions = append(instructions, opcodes.ICONST_4)
	case 5:
		instructions = append(instructions, opcodes.ICONST_5)
	default:
		instructions = append(instructions, opcodes.LDC)
		instructions = jvmio.AppendPaddedBytes(instructions, pool.FindIntConstant(i.Constant), 1)
	}

	return instructions
}

func (i IntConstant) FillConstantsPool(pool *constantpool.ConstantPool) {
	if i.requiresPoolEntry() {
		pool.AddIntConstant(i.Constant)
	}
}

func (i IntConstant) MaxStack() int {
	return 1
}

func (i IntConstant) requiresPoolEntry() bool {
	return i.Constant > 5 || i.Constant < -1
}

func NewIntConstant(constant int) IntConstant {
	return IntConstant{Constant: constant}
}

type StringConstant struct {
	Constant string
}

func (s StringConstant) GetInstructions(stack *Stack, pool *constantpool.ConstantPool) []byte {
	instructions := []byte{opcodes.LDC}
	return jvmio.AppendPaddedBytes(instructions, pool.FindStringConstant(s.Constant), 1)
}

func (s StringConstant) FillConstantsPool(pool *constantpool.ConstantPool) {
	pool.AddStringConstant(s.Constant)
}

func (s StringConstant) MaxStack() int {
	return 1
}

func NewStringConstant(constant string) StringConstant {
	return StringConstant{Constant: constant}
}
