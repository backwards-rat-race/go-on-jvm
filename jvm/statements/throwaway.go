package statements

import (
	"go-on-jvm/jvm/constantpool"
	jvmio "go-on-jvm/jvm/io"
	"go-on-jvm/jvm/opcodes"
)

type ThrowAway struct {
	Statement Statement
}

func (t ThrowAway) GetInstructions(stack *Stack, pool *constantpool.ConstantPool) []byte {
	instructions := t.Statement.GetInstructions(stack, pool)
	instructions = jvmio.AppendPaddedBytes(instructions, opcodes.POP, 1)
	return instructions
}

func (t ThrowAway) FillConstantsPool(pool *constantpool.ConstantPool) {
	t.Statement.FillConstantsPool(pool)
}

func (t ThrowAway) MaxStack() uint {
	return t.Statement.MaxStack()
}

func NewThrowAway(statement Statement) ThrowAway {
	return ThrowAway{Statement: statement}
}
