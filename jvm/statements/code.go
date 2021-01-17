package statements

import (
	"bytes"
	"go-on-jvm/jvm/constantpool"
	jvmio "go-on-jvm/jvm/io"
	"io"
)

type codeAttributeSerialiser struct {
	Stack Stack
	pool  *constantpool.ConstantPool
}

func newCodeAttributeSerialiser(stack Stack, pool *constantpool.ConstantPool) jvmio.Serialisable {
	return codeAttributeSerialiser{stack, pool}
}

func (c codeAttributeSerialiser) Write(w io.Writer) error {
	// u2 attribute_name_index;
	err := jvmio.WritePaddedBytes(w, c.pool.FindUTF8Item(CodeAttribute), 2)
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
	// u2 max_stack
	err := jvmio.WritePaddedBytes(w, c.Stack.MaxSize(), 2)
	if err != nil {
		return err
	}

	// u2 max_locals
	err = jvmio.WritePaddedBytes(w, c.Stack.MaxLocals(), 2)
	if err != nil {
		return err
	}

	var buffer bytes.Buffer
	err = c.writeCodeData(&buffer)
	if err != nil {
		return err
	}

	// u4 code_length
	err = jvmio.WritePaddedBytes(w, buffer.Len(), 4)
	if err != nil {
		return err
	}

	// u1 code[code_length];
	_, err = io.Copy(w, &buffer)
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

func (c codeAttributeSerialiser) writeCodeData(w io.Writer) error {
	for _, statement := range c.Stack.Statements {
		err := statement.NewSerialiser(c.Stack, c.pool).Write(w)
		if err != nil {
			return err
		}
	}
	return nil
}
