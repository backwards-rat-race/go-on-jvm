package statements

import (
	"go-on-jvm/jvm/constantpool"
	"go-on-jvm/jvm/opcodes"
)

type If struct {
	Condition Comparison
	Success   Block
	Failure   Block
}

func (i If) GetInstructions(stack *Stack, pool *constantpool.ConstantPool) []byte {

	// TODO write instructions backwards to know jump lengths?

	if i.Success.Empty() {
		// Panic?
		return nil
	}

	nestedInstructions := i.Success.GetInstructions(stack, pool)
	jumpTo := uint(opcodes.IfSize)

	if i.Failure.Empty() {
		jumpTo += uint(len(nestedInstructions))
	} else {
		failureInstructions := i.Failure.GetInstructions(stack, pool)

		// Now we know how long the failure branch is we can add the goto instruction
		// at the end of the success block (the instruction after failure
		gotoJump := len(failureInstructions) + opcodes.GotoSize
		nestedInstructions = append(nestedInstructions, opcodes.GetGotoInstructionI(gotoJump)...)
		jumpTo += uint(len(nestedInstructions))
		nestedInstructions = append(nestedInstructions, failureInstructions...)
	}

	instructions := i.Condition.GetInstructions(jumpTo, stack, pool)

	instructions = append(instructions, nestedInstructions...)

	return instructions
}

func (i If) FillConstantsPool(pool *constantpool.ConstantPool) {
	i.Condition.FillConstantsPool(pool)
	i.Success.FillConstantsPool(pool)
	i.Failure.FillConstantsPool(pool)
}

func (i If) MaxStack() uint {
	return iMax(i.Failure.MaxStack(), i.Success.MaxStack())
}

func NewIf(comparison Comparison) If {
	return If{Condition: comparison}
}
