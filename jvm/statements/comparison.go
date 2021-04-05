package statements

import (
	"fmt"
	"go-on-jvm/jvm/constantpool"
	jvmio "go-on-jvm/jvm/io"
	"go-on-jvm/jvm/opcodes"
)

type ComparisonType int

const (
	IsNull ComparisonType = iota
	IsNotNull
	IsEqual
	IsNotEqual
	IsLessThan
	IsLessThanOrEqual
	IsGreaterThan
	IsGreaterThanOrEqual
)

var CompareNothing = Comparison{}

func (c ComparisonType) opcode() uint {
	switch c {
	case IsNull:
		return opcodes.IFNULL
	case IsNotNull:
		return opcodes.IFNONNULL
	case IsEqual:
		return opcodes.IFEQ
	case IsNotEqual:
		return opcodes.IFNE
	case IsLessThan:
		return opcodes.IFLT
	case IsLessThanOrEqual:
		return opcodes.IFLE
	case IsGreaterThan:
		return opcodes.IFGT
	case IsGreaterThanOrEqual:
		return opcodes.IFGE
	default:
		panic(fmt.Errorf("unknown if comparison: %d", c))
	}
}

func (c ComparisonType) Inverse() ComparisonType {
	switch c {
	case IsNull:
		return IsNotNull
	case IsNotNull:
		return IsNull
	case IsEqual:
		return IsNotEqual
	case IsNotEqual:
		return IsEqual
	case IsLessThan:
		return IsGreaterThanOrEqual
	case IsLessThanOrEqual:
		return IsGreaterThan
	case IsGreaterThan:
		return IsLessThanOrEqual
	case IsGreaterThanOrEqual:
		return IsLessThan
	default:
		panic(fmt.Errorf("unknown comparison: %d", c))
	}
}

func (c ComparisonType) New(statement Statement) Comparison {
	return Comparison{Statement: statement}
}

type Comparison struct {
	Type      ComparisonType
	Statement Statement
}

func (c Comparison) GetInstructions(jump uint, stack *Stack, pool *constantpool.ConstantPool) []byte {
	if c == CompareNothing {
		return nil
	}

	instructions := c.Statement.GetInstructions(stack, pool)

	// Then write the if operation. We inverse this as the if acts as a jump. So if we're replicating an if null.
	// The operation is 'ifnotnull #34' for example. If we have both a success and failure block it doesn't make
	// much difference. But if we only have a success block then the operation needs to be inverse. So we just do
	// it everytime.
	instructions = jvmio.AppendPaddedBytes(instructions, c.Type.Inverse().opcode(), 1)
	instructions = jvmio.AppendPaddedBytes(instructions, jump, 2)
	return instructions
}

func (c Comparison) FillConstantsPool(pool *constantpool.ConstantPool) {
	c.Statement.FillConstantsPool(pool)
}

func (c Comparison) MaxStack() uint {
	return c.Statement.MaxStack()
}
