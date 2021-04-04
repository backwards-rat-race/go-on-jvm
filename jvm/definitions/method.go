package definitions

import (
	"bytes"
	"go-on-jvm/jvm/constantpool"
	jvmio "go-on-jvm/jvm/io"
	"go-on-jvm/jvm/statements"
	jvmtypes "go-on-jvm/jvm/types"
	"io"
)

type Method struct {
	statements.Block
	Name       string
	ReturnType jvmtypes.TypeReference
	Arguments  []statements.Variable
	Access     []AccessModifier
}

func NewMethod(name string, access ...AccessModifier) Method {
	return Method{
		Name:   name,
		Access: access,
	}
}

func (m *Method) AddArgument(argument statements.Variable) {
	m.Arguments = append(m.Arguments, argument)
}

func (m *Method) Type() jvmtypes.MethodType {
	var argTypes []jvmtypes.TypeReference
	for _, argument := range m.Arguments {
		if argument == statements.SelfReferenceVariable {
			continue
		}
		argTypes = append(argTypes, argument.Type)
	}
	return jvmtypes.MethodType{ReturnType: m.ReturnType, Arguments: argTypes}
}

func (m Method) HasAccessModifier(modifier AccessModifier) bool {
	for _, access := range m.Access {
		if access == modifier {
			return true
		}
	}
	return false
}

func (m Method) CreateStack() *statements.Stack {
	var arguments []statements.Variable

	if m.HasAccessModifier(Static) {
		arguments = m.Arguments
	} else {
		arguments = []statements.Variable{statements.SelfReferenceVariable}
		arguments = append(arguments, m.Arguments...)
	}

	stack := statements.NewStack(arguments...)

	if !m.HasAccessModifier(Static) {
		m.AddArgument(statements.SelfReferenceVariable)
	}

	return stack
}

func (m Method) fillConstantsPool(pool *constantpool.ConstantPool) {
	pool.AddUTF8(m.Name)
	pool.AddUTF8(m.Type().Descriptor())
	m.Block.FillConstantsPool(pool)
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
	err = jvmio.WritePaddedBytes(w, m.Pool.FindUTF8Item(m.Type().Descriptor()), 2)
	if err != nil {
		return err
	}

	// TODO: Cleanup. Attribute writer?
	if m.Empty() {
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
		err = m.NewCodeAttributeSerialiser(m.Pool).Write(w)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m methodSerialiser) NewCodeAttributeSerialiser(pool *constantpool.ConstantPool) jvmio.Serialisable {
	return newCodeAttributeSerialiser(m.CreateStack(), m.Method, pool)
}

type codeAttributeSerialiser struct {
	Stack  *statements.Stack
	Method Method
	pool   *constantpool.ConstantPool
}

func newCodeAttributeSerialiser(stack *statements.Stack, method Method, pool *constantpool.ConstantPool) jvmio.Serialisable {
	return codeAttributeSerialiser{stack, method, pool}
}

func (c codeAttributeSerialiser) Write(w io.Writer) error {
	// u2 attribute_name_index;
	err := jvmio.WritePaddedBytes(w, c.pool.FindUTF8Item(statements.CodeAttribute), 2)
	if err != nil {
		return err
	}

	// We now write the data after the length indicator to a temporary
	// buffer to allow us to know the length before writing to the
	// 'real' writer
	var buffer bytes.Buffer
	err = c.writeAttributeData(&buffer)
	if err != nil {
		return err
	}

	// Written internal data. Now we know the length
	// u4 attribute_length;
	err = jvmio.WritePaddedBytes(w, buffer.Len(), 4)
	if err != nil {
		return err
	}

	_, err = io.Copy(w, &buffer)
	if err != nil {
		return err
	}

	return nil
}

func (c codeAttributeSerialiser) writeAttributeData(w io.Writer) error {

	stack := c.Method.CreateStack()
	instructions := c.Method.GetInstructions(0, stack, c.pool)

	// u2 max_stack
	err := jvmio.WritePaddedBytes(w, stack.MaxSize(), 2)
	if err != nil {
		return err
	}

	// u2 max_locals
	err = jvmio.WritePaddedBytes(w, stack.MaxLocals(), 2)
	if err != nil {
		return err
	}

	// u4 code_length
	err = jvmio.WritePaddedBytes(w, len(instructions), 4)
	if err != nil {
		return err
	}

	// u1 code[code_length];
	_, err = w.Write(instructions)
	if err != nil {
		return err

	}

	// u2 exception_table_length
	_, err = w.Write([]byte{0x00, 0x00})
	if err != nil {
		return err
	}

	// u2 attributes_count
	_, err = w.Write([]byte{0x00, 0x00})
	if err != nil {
		return err
	}

	return nil
}
