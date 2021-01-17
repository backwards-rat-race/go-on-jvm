package statements

import (
	"go-on-jvm/jvm/constantpool"
	jvmio "go-on-jvm/jvm/io"
	"go-on-jvm/jvm/opcodes"
	"io"
)

type MethodReference struct {
	Class string
	Name  string
	Type  string
}

func NewMethodReference(class string, name string, returnType string, args ...string) MethodReference {
	return MethodReference{
		Class: class,
		Name:  name,
		//Type:  buildMethodTypeDescriptor(returnType, args...),
	}
}

type Invocation struct {
	MethodReference MethodReference
	Static          bool
	Vars            []Variable
}

func NewInvocation(method MethodReference, vars ...Variable) Invocation {
	return Invocation{
		MethodReference: method,
		Vars:            vars,
	}
}

func NewStaticInvocation(method MethodReference, vars ...Variable) Invocation {
	return Invocation{
		MethodReference: method,
		Static:          true,
		Vars:            vars,
	}
}

func (i Invocation) NewSerialiser(stack Stack, pool *constantpool.ConstantPool) jvmio.Serialisable {
	return newInvocationStatementSerialiser(i, stack, pool)
}

func (i Invocation) Variables() []Variable {
	return i.Vars
}

func (i Invocation) fillConstantsPool(pool *constantpool.ConstantPool) {
	pool.AddMethodReference(i.MethodReference.Class, i.MethodReference.Name, i.MethodReference.Type)
}

type invocationStatementSerialiser struct {
	Invocation
	Stack Stack
	Pool  *constantpool.ConstantPool
}

func newInvocationStatementSerialiser(statement Invocation, stack Stack, pool *constantpool.ConstantPool) *invocationStatementSerialiser {
	return &invocationStatementSerialiser{statement, stack, pool}
}

func (i invocationStatementSerialiser) Write(w io.Writer) error {
	for _, variable := range i.Variables() {
		index := i.Stack.Index(variable)
		op := opcodes.NewALoad(index)
		err := jvmio.WritePaddedBytes(w, op, 1)
		if err != nil {
			return err
		}
	}

	err := jvmio.WritePaddedBytes(w, i.opcode(), 1)
	if err != nil {
		return err
	}

	index := i.Pool.FindMethodReference(i.MethodReference.Class, i.MethodReference.Name, i.MethodReference.Type)
	err = jvmio.WritePaddedBytes(w, index, 2)
	if err != nil {
		return err
	}

	return nil
}

func (i invocationStatementSerialiser) opcode() int {
	if i.Static {
		return opcodes.INVOKESTATIC
	} else {
		return opcodes.INVOKESPECIAL
	}
}
