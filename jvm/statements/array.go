package statements

import (
	"go-on-jvm/jvm/constantpool"
	jvmio "go-on-jvm/jvm/io"
	"go-on-jvm/jvm/opcodes"
)

type ArrayLen struct {
	Array Statement
}

func (a ArrayLen) GetInstructions(stack *Stack, pool *constantpool.ConstantPool) []byte {
	instructions := a.Array.GetInstructions(stack, pool)
	instructions = jvmio.AppendPaddedBytes(instructions, opcodes.ARRAYLENGTH, 1)
	return instructions
}

func (a ArrayLen) FillConstantsPool(pool *constantpool.ConstantPool) {
	a.Array.FillConstantsPool(pool)
}

func (a ArrayLen) MaxStack() uint {
	if a.Array == nil {
		return 0
	} else {
		return a.Array.MaxStack()
	}
}

func NewArrayLen(array Statement) ArrayLen {
	return ArrayLen{array}
}

type ArrayGet struct {
	Array Statement
	Index Statement
}

func (a ArrayGet) GetInstructions(stack *Stack, pool *constantpool.ConstantPool) []byte {
	instructions := a.Array.GetInstructions(stack, pool)
	instructions = append(instructions, a.Index.GetInstructions(stack, pool)...)
	instructions = jvmio.AppendPaddedBytes(instructions, opcodes.AALOAD, 1)
	return instructions
}

func (a ArrayGet) FillConstantsPool(pool *constantpool.ConstantPool) {
	a.Array.FillConstantsPool(pool)
	a.Index.FillConstantsPool(pool)
}

func (a ArrayGet) MaxStack() uint {
	var max uint = 2
	max = iMax(max, a.Array.MaxStack())
	max = iMax(max, a.Index.MaxStack())
	return max
}

func NewArrayGet(array Statement, index Statement) ArrayGet {
	return ArrayGet{Array: array, Index: index}
}
