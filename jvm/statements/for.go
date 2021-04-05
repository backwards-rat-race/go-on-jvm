package statements

import (
	"go-on-jvm/jvm/constantpool"
)

type ForLoop struct {
	Block       Block
	Condition   Comparison
	Declaration Statement
	Increment   Statement
}

func (f ForLoop) GetInstructions(stack *Stack, pool *constantpool.ConstantPool) []byte {
	// TODO
	return nil
}

func (f ForLoop) FillConstantsPool(pool *constantpool.ConstantPool) {
	if f.Declaration != nil {
		f.Declaration.FillConstantsPool(pool)
	}

	f.Condition.FillConstantsPool(pool)

	if f.Increment != nil {
		f.Increment.FillConstantsPool(pool)
	}
}

func (f ForLoop) MaxStack() int {
	max := 1

	if f.Declaration != nil {
		max = iMax(max, f.Declaration.MaxStack())
	}

	max = iMax(max, f.Condition.MaxStack())

	if f.Increment != nil {
		max = iMax(max, f.Increment.MaxStack())
	}

	return max
}

func NewForLoop(block Block, declaration Statement, condition Comparison, increment Statement) ForLoop {
	return ForLoop{Block: block, Declaration: declaration, Condition: condition, Increment: increment}
}
