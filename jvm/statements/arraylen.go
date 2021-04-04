package statements

import (
	"go-on-jvm/jvm/constantpool"
	jvmio "go-on-jvm/jvm/io"
	"go-on-jvm/jvm/opcodes"
)

type ArrayLen struct {
	Array Statement
}

func (a ArrayLen) GetInstructions(writeIndex int, stack *Stack, pool *constantpool.ConstantPool) []byte {
	instructions := a.Array.GetInstructions(writeIndex, stack, pool)
	instructions = jvmio.AppendPaddedBytes(instructions, opcodes.ARRAYLENGTH, 1)
	return instructions
}

func (a ArrayLen) FillConstantsPool(pool *constantpool.ConstantPool) {
	a.Array.FillConstantsPool(pool)
}

func NewArrayLen(array Statement) ArrayLen {
	return ArrayLen{array}
}
