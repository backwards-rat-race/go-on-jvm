package constantpool

import (
	jvmio "go-on-jvm/jvm/io"
	"io"
)

type NameAndType struct {
	Name string
	Type string
}

func newNameAndType(name string, typeDescriptor string) NameAndType {
	return NameAndType{name, typeDescriptor}
}

func (n NameAndType) isTag(tag ConstantPoolTag) bool {
	return tag == NameAndTypeDescriptor
}

func (n NameAndType) write(w io.Writer, constantPool ConstantPool, _ int) error {
	err := NameAndTypeDescriptor.write(w)
	if err != nil {
		return err
	}

	err = jvmio.WritePaddedBytes(w, constantPool.FindUTF8Item(n.Name), 2)
	if err != nil {
		return err
	}

	err = jvmio.WritePaddedBytes(w, constantPool.FindUTF8Item(n.Type), 2)
	if err != nil {
		return err
	}

	return nil
}

func isNameAndType(c constantPoolItem, name string, typeDescriptor string) bool {
	return c.isTag(NameAndTypeDescriptor) &&
		c.(NameAndType).Name == name &&
		c.(NameAndType).Type == typeDescriptor
}
