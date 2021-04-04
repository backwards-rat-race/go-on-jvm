package statements

import (
	"fmt"
	"go-on-jvm/jvm/constantpool"
	jvmio "go-on-jvm/jvm/io"
	"go-on-jvm/jvm/opcodes"
)

type ReturnType int

const (
	ReturnVoid ReturnType = iota
	ReturnReference
	ReturnInt
	ReturnLong
	ReturnFloat
	ReturnDouble
)

func (r ReturnType) opcode() int {
	switch r {
	case ReturnVoid:
		return opcodes.RETURN
	case ReturnReference:
		return opcodes.ARETURN
	case ReturnInt:
		return opcodes.IRETURN
	case ReturnLong:
		return opcodes.LRETURN
	case ReturnFloat:
		return opcodes.FRETURN
	case ReturnDouble:
		return opcodes.DRETURN
	default:
		panic(fmt.Errorf("unknown return type: %d", r))
	}
}

type Return struct {
	Type      ReturnType
	Statement Statement
}

func (r Return) GetInstructions(writeIndex int, stack *Stack, pool *constantpool.ConstantPool) []byte {
	var instructions []byte

	if r.Type == ReturnVoid {
		instructions = make([]byte, 0)
	} else {
		instructions = r.Statement.GetInstructions(writeIndex, stack, pool)
	}

	return jvmio.AppendPaddedBytes(instructions, r.Type.opcode(), 1)
}

func (r Return) FillConstantsPool(pool *constantpool.ConstantPool) {
	if r.Statement != nil {
		r.Statement.FillConstantsPool(pool)
	}
}

func NewVoidReturn() Return {
	return Return{Type: ReturnVoid}
}

func NewReturn(returnType ReturnType, statement Statement) Return {
	return Return{Type: returnType, Statement: statement}
}
