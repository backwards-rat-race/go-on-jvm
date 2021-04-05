package statements

import (
	"go-on-jvm/jvm/constantpool"
)

const CodeAttribute = "Code"

type Block struct {
	Statements []Statement
}

func (b Block) GetInstructions(stack *Stack, pool *constantpool.ConstantPool) []byte {
	instructions := make([]byte, 0)

	for _, statement := range b.Statements {
		instructions = append(instructions, statement.GetInstructions(stack, pool)...)
	}

	return instructions
}

func (b Block) MaxStack() int {
	max := 0

	for _, statement := range b.Statements {
		max = iMax(max, statement.MaxStack())
	}

	return max
}

func (b *Block) AddStatement(statement Statement) {
	b.Statements = append(b.Statements, statement)
}

func (b Block) Empty() bool {
	return len(b.Statements) == 0
}

func (b Block) FillConstantsPool(pool *constantpool.ConstantPool) {
	if b.Empty() {
		return
	}

	pool.AddUTF8(CodeAttribute)

	for _, statement := range b.Statements {
		statement.FillConstantsPool(pool)
	}
}

func NewBlock() Block {
	return Block{}
}

func iMax(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
