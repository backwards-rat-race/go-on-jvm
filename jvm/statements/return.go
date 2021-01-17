package statements

import (
	"fmt"
	"go-on-jvm/jvm/constantpool"
	jvmio "go-on-jvm/jvm/io"
	"go-on-jvm/jvm/opcodes"
	"io"
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

func (r ReturnType) NewSerialiser(_ Stack, _ *constantpool.ConstantPool) jvmio.Serialisable {
	return newReturnSerialiser(r)
}

func (r ReturnType) Variables() []Variable {
	return nil
}

func (r ReturnType) fillConstantsPool(_ *constantpool.ConstantPool) {
}

type returnTypeSerialiser struct {
	ReturnType
}

func newReturnSerialiser(returnType ReturnType) *returnTypeSerialiser {
	return &returnTypeSerialiser{returnType}
}

func (r returnTypeSerialiser) Write(w io.Writer) error {
	return jvmio.WritePaddedBytes(w, r.opcode(), 1)
}

func (r returnTypeSerialiser) opcode() int {
	switch r.ReturnType {
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
		panic(fmt.Errorf("unknown return type: %d", r.ReturnType))
	}
}
