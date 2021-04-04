package statements

import (
	"go-on-jvm/jvm/constantpool"
)

const CodeAttribute = "Code"

type Block struct {
	Statements []Statement
}

func (b Block) GetInstructions(writeIndex int, stack *Stack, pool *constantpool.ConstantPool) []byte {
	instructions := make([]byte, 0)

	for _, statement := range b.Statements {
		index := writeIndex + len(instructions)
		instructions = append(instructions, statement.GetInstructions(index, stack, pool)...)
	}

	return instructions
}

func NewBlock() Block {
	return Block{}
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
