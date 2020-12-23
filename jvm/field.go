package jvm

import (
	"go-on-jvm/jvm/constants"
	jvmio "go-on-jvm/jvm/io"
	"io"
)

type Field struct {
	Name   string
	Access []AccessModifier
	Type   string
}

func NewField(name string, typeDescriptor string, access ...AccessModifier) Field {
	return Field{name, access, typeDescriptor}
}

func (f Field) fillConstantsPool(pool *constants.ConstantPool) {
	pool.AddUTF8(f.Name)
	pool.AddUTF8(f.Type)
}

type fieldCompiler struct {
	Field
	Pool *constants.ConstantPool
}

func newFieldCompiler(field Field, pool *constants.ConstantPool) *fieldCompiler {
	return &fieldCompiler{field, pool}
}

func (f *fieldCompiler) Compile(w io.Writer) error {
	err := writeAccessModifier(w, f.Access)
	if err != nil {
		return err
	}

	err = jvmio.WritePaddedBytes(w, f.Pool.FindUTF8Item(f.Name), 2)
	if err != nil {
		return err
	}

	err = jvmio.WritePaddedBytes(w, f.Pool.FindUTF8Item(f.Type), 2)
	if err != nil {
		return err
	}

	// Attributes count
	_, err = w.Write([]byte{0x00, 0x00})
	if err != nil {
		return err
	}

	return nil
}
