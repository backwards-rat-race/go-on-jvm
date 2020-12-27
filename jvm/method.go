package jvm

import (
	"go-on-jvm/jvm/constantpool"
	jvmio "go-on-jvm/jvm/io"
	"io"
)

type Method struct {
	Name       string
	Access     []AccessModifier
	Type       string
	Statements []Statement
}

func NewMethod(name string, access ...AccessModifier) Method {
	return Method{name, access, "", make([]Statement, 0)}
}

func (m *Method) WithTypeDescriptor(returnType string, argTypes ...string) {
	typeDescriptor := "("

	for _, argType := range argTypes {
		typeDescriptor += argType
		typeDescriptor += ";"
	}

	typeDescriptor += ")"
	typeDescriptor += returnType

	m.Type = typeDescriptor
}

func (m Method) fillConstantsPool(pool *constantpool.ConstantPool) {

	pool.AddUTF8(m.Name)
	pool.AddUTF8(m.Type)

	for _, statement := range m.Statements {
		statement.fillConstantsPool(pool)
	}
}

type methodCompiler struct {
	Method
	Pool *constantpool.ConstantPool
}

func newMethodCompiler(method Method, pool *constantpool.ConstantPool) *methodCompiler {
	return &methodCompiler{method, pool}
}

func (m methodCompiler) Compile(w io.Writer) error {
	err := writeAccessModifier(w, m.Access)
	if err != nil {
		return err
	}

	err = jvmio.WritePaddedBytes(w, m.Pool.FindUTF8Item(m.Name), 2)
	if err != nil {
		return err
	}

	err = jvmio.WritePaddedBytes(w, m.Pool.FindUTF8Item(m.Type), 2)
	if err != nil {
		return err
	}

	_, err = w.Write([]byte{0x00, 0x00})
	if err != nil {
		return err
	}

	return nil
}
