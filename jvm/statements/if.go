package statements

import (
	"fmt"
	"go-on-jvm/jvm/constantpool"
	jvmio "go-on-jvm/jvm/io"
	"go-on-jvm/jvm/opcodes"
)

type IfType int

const (
	IfNull IfType = iota
	IfNotNull
	IfEqual
	IfNotEqual
	IfLessThan
	IfLessThanOrEqual
	IfGreaterThan
	IfGreaterThanOrEqual
)

func (i IfType) New(condition Statement) If {
	return If{Type: i, Condition: condition}
}

func (i IfType) opcode() int {
	switch i {
	case IfNull:
		return opcodes.IFNULL
	case IfNotNull:
		return opcodes.IFNONNULL
	case IfEqual:
		return opcodes.IFEQ
	case IfNotEqual:
		return opcodes.IFNE
	case IfLessThan:
		return opcodes.IFLT
	case IfLessThanOrEqual:
		return opcodes.IFLE
	case IfGreaterThan:
		return opcodes.IFGT
	case IfGreaterThanOrEqual:
		return opcodes.IFGE
	default:
		panic(fmt.Errorf("unknown if comparison: %d", i))
	}
}

func (i IfType) Inverse() IfType {
	switch i {
	case IfNull:
		return IfNotNull
	case IfNotNull:
		return IfNull
	case IfEqual:
		return IfNotEqual
	case IfNotEqual:
		return IfEqual
	case IfLessThan:
		return IfGreaterThanOrEqual
	case IfLessThanOrEqual:
		return IfGreaterThan
	case IfGreaterThan:
		return IfLessThanOrEqual
	case IfGreaterThanOrEqual:
		return IfLessThan
	default:
		panic(fmt.Errorf("unknown if comparison: %d", i))
	}
}

type If struct {
	Type      IfType
	Condition Statement
	Success   Block
	Failure   Block
}

func (i If) GetInstructions(writeIndex int, stack *Stack, pool *constantpool.ConstantPool) []byte {
	instructions := make([]byte, 0)

	if i.Success.Empty() {
		// Panic?
		return instructions
	}

	// First load the condition onto the stack
	instructions = append(instructions, i.Condition.GetInstructions(writeIndex, stack, pool)...)
	writeIndex += len(instructions)

	// We can't write it yet as we don't know where we're jumping
	// off too, but we know it's going to be 3 bytes.
	writeIndex += 3

	nestedInstructions := i.Success.GetInstructions(writeIndex, stack, pool)
	writeIndex += len(nestedInstructions)

	if !i.Failure.Empty() {
		// We're going to require a goto at the end of the success block. We can't write it yet
		// as we don't know where we're jumping too. We know the size will be 3 though
		// FIXME support goto_w which is 5 bytes?
		writeIndex += 3

		failureInstructions := i.Failure.GetInstructions(writeIndex, stack, pool)
		writeIndex += len(failureInstructions)

		// Now we know how long the failure branch is we can add the goto instruction
		// at the end of the success block
		nestedInstructions = append(nestedInstructions, opcodes.GetGotoInstruction(writeIndex)...)
		nestedInstructions = append(nestedInstructions, failureInstructions...)
	}

	// Then write the if operation. We inverse this as the if acts as a jump. So if we're replicating an if null.
	// The operation is 'ifnotnull #34' for example. If we have both a success and failure block it doesn't make
	// much difference. But if we only have a success block then the operation needs to be inverse. So we just do
	// it everytime.
	instructions = jvmio.AppendPaddedBytes(instructions, i.Type.Inverse().opcode(), 1)

	// Where we want to jump to
	instructions = jvmio.AppendPaddedBytes(instructions, writeIndex, 2)

	instructions = append(instructions, nestedInstructions...)

	return instructions
}

func (i If) FillConstantsPool(pool *constantpool.ConstantPool) {
	i.Condition.FillConstantsPool(pool)
	i.Success.FillConstantsPool(pool)
	i.Failure.FillConstantsPool(pool)
}
