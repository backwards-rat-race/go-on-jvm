package definitions

import (
	"go-on-jvm/jvm/constantpool"
	jvmio "go-on-jvm/jvm/io"
	"go-on-jvm/jvm/statements"
	"io"
)

type Method struct {
	Name   string
	Type   string
	Access []AccessModifier
	Stack  statements.Stack
}

func NewMethod(name string, access ...AccessModifier) Method {
	return Method{
		Name:   name,
		Access: access,
	}
}

func (m *Method) WithTypeDescriptor(returnType string, argTypes ...string) {
	m.Type = buildMethodTypeDescriptor(returnType, argTypes...)
}

func (m *Method) AddStatement(statement statements.Statement) {
	m.Stack.Statements = append(m.Stack.Statements, statement)
}

func (m *Method) SetStatements(statement ...statements.Statement) {
	m.Stack.Statements = statement
}

func (m Method) fillConstantsPool(pool *constantpool.ConstantPool) {

	pool.AddUTF8(m.Name)
	pool.AddUTF8(m.Type)
	m.Stack.FillConstantsPool(pool)
}

type methodSerialiser struct {
	Method
	Pool *constantpool.ConstantPool
}

func newMethodSerialiser(method Method, pool *constantpool.ConstantPool) *methodSerialiser {
	return &methodSerialiser{method, pool}
}

func (m methodSerialiser) Write(w io.Writer) error {
	// u2 access_flags;
	err := writeAccessModifier(w, m.Access)
	if err != nil {
		return err
	}

	// u2 name_index;
	err = jvmio.WritePaddedBytes(w, m.Pool.FindUTF8Item(m.Name), 2)
	if err != nil {
		return err
	}

	// u2 descriptor_index;
	err = jvmio.WritePaddedBytes(w, m.Pool.FindUTF8Item(m.Type), 2)
	if err != nil {
		return err
	}

	// TODO: Cleanup. Attribute writer?
	if m.Stack.Empty() {
		// u2 attributes_count;
		err = jvmio.WritePaddedBytes(w, 0, 2)
		if err != nil {
			return err
		}
	} else {
		// u2 attributes_count;
		err = jvmio.WritePaddedBytes(w, 1, 2)
		if err != nil {
			return err
		}

		// attribute_info attributes[attributes_count];
		err = m.Stack.NewCodeAttributeSerialiser(m.Pool).Write(w)
		if err != nil {
			return err
		}
	}

	return nil
}

func buildMethodTypeDescriptor(returnType string, argTypes ...string) string {
	typeDescriptor := "("

	for _, argType := range argTypes {
		typeDescriptor += argType
	}

	typeDescriptor += ")"
	typeDescriptor += returnType

	return typeDescriptor
}
