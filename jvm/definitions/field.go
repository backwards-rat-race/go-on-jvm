package definitions

import (
	"go-on-jvm/jvm/constantpool"
	jvmio "go-on-jvm/jvm/io"
	jvmtypes "go-on-jvm/jvm/types"
	"io"
)

type Field struct {
	Name   string
	Access []AccessModifier
	Type   jvmtypes.TypeReference
}

func NewField(name string, typeDescriptor jvmtypes.TypeReference, access ...AccessModifier) Field {
	return Field{name, access, typeDescriptor}
}

func (f *Field) WithAccess(modifier ...AccessModifier) {
	f.Access = modifier
}

func (f Field) fillConstantsPool(pool *constantpool.ConstantPool) {
	pool.AddUTF8(f.Name)
	pool.AddUTF8(f.Type.Jvm())
}

type fieldSerialiser struct {
	Field
	Pool *constantpool.ConstantPool
}

func newFieldSerialiser(field Field, pool *constantpool.ConstantPool) *fieldSerialiser {
	return &fieldSerialiser{field, pool}
}

func (f *fieldSerialiser) Write(w io.Writer) error {
	err := writeAccessModifier(w, f.Access)
	if err != nil {
		return err
	}

	err = jvmio.WritePaddedBytes(w, f.Pool.FindUTF8Item(f.Name), 2)
	if err != nil {
		return err
	}

	err = jvmio.WritePaddedBytes(w, f.Pool.FindUTF8Item(f.Type.Jvm()), 2)
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
