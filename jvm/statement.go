package jvm

import (
	"go-on-jvm/jvm/constantpool"
	jvmio "go-on-jvm/jvm/io"
	"go-on-jvm/jvm/opcodes"
	"io"
)

type Statement interface {
	NewSerialiser(stack Stack, pool *constantpool.ConstantPool) Serialisable
	Variables() []Variable
	fillConstantsPool(pool *constantpool.ConstantPool)
}

type ReturnStatement struct{}

func (r ReturnStatement) NewSerialiser(_ Stack, _ *constantpool.ConstantPool) Serialisable {
	return newReturnStatementSerialiser(r)
}

func (r ReturnStatement) Variables() []Variable {
	return nil
}

func (r ReturnStatement) fillConstantsPool(pool *constantpool.ConstantPool) {
}

func NewReturnStatement() ReturnStatement {
	return ReturnStatement{}
}

type returnStatementSerialiser struct{}

func newReturnStatementSerialiser(_ ReturnStatement) *returnStatementSerialiser {
	return &returnStatementSerialiser{}
}

func (r returnStatementSerialiser) Write(w io.Writer) error {
	return jvmio.WritePaddedBytes(w, opcodes.RETURN, 1)
}

type InvocationStatement struct {
	MethodReference MethodReference
	Static          bool
	Vars            []Variable
}

func NewInvocationStatement(method MethodReference, vars ...Variable) InvocationStatement {
	return InvocationStatement{
		MethodReference: method,
		Vars:            vars,
	}
}

func NewStaticInvocationStatement(method MethodReference, vars ...Variable) InvocationStatement {
	return InvocationStatement{
		MethodReference: method,
		Static:          true,
		Vars:            vars,
	}
}

func (i InvocationStatement) NewSerialiser(stack Stack, pool *constantpool.ConstantPool) Serialisable {
	return newInvocationStatementSerialiser(i, stack, pool)
}

func (i InvocationStatement) Variables() []Variable {
	return i.Vars
}

func (i InvocationStatement) fillConstantsPool(pool *constantpool.ConstantPool) {
	pool.AddMethodReference(i.MethodReference.Class, i.MethodReference.Name, i.MethodReference.Type)
}

type invocationStatementSerialiser struct {
	InvocationStatement
	Stack Stack
	Pool  *constantpool.ConstantPool
}

func newInvocationStatementSerialiser(statement InvocationStatement, stack Stack, pool *constantpool.ConstantPool) *invocationStatementSerialiser {
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

	err := jvmio.WritePaddedBytes(w, opcodes.INVOKESPECIAL, 1)
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
