package statements

import (
	"go-on-jvm/jvm/constantpool"
	"go-on-jvm/jvm/opcodes"
)

type ForLoop struct {
	Block       Block
	Condition   Comparison
	Declaration Statement
	Increment   Statement
}

func (f ForLoop) GetInstructions(stack *Stack, pool *constantpool.ConstantPool) []byte {

	// First is the declaration
	// Second is the comparison (this is where the loop jumps back to)
	// Then is the block
	// Then the increment
	// And finally a goto to jump back to the comparison

	var instructions []byte
	if f.Declaration != nil {
		instructions = f.Declaration.GetInstructions(stack, pool)
	}

	nestedInstructions := f.Block.GetInstructions(stack, pool)

	if f.Increment != nil {
		nestedInstructions = append(nestedInstructions, f.Increment.GetInstructions(stack, pool)...)
	}

	nestedInstructions = append(nestedInstructions, opcodes.GetGotoInstructionI(-len(nestedInstructions))...)

	instructions = append(instructions, f.Condition.GetInstructions(uint(len(instructions)), stack, pool)...)
	instructions = append(instructions, nestedInstructions...)

	return instructions
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

func (f ForLoop) MaxStack() uint {
	var max uint = 1

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
