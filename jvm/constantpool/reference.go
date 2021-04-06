package constantpool

import (
	jvmio "go-on-jvm/jvm/io"
	"io"
)

type ReferenceItem struct {
	Tag   ConstantPoolTag
	Class string
	Name  string
	Type  string
}

func newMethodReference(class, name, typeDescriptor string) ReferenceItem {
	return ReferenceItem{MethodReference, class, name, typeDescriptor}
}

func newFieldReference(class, name, typeDescriptor string) ReferenceItem {
	return ReferenceItem{FieldReference, class, name, typeDescriptor}
}

func (m ReferenceItem) isTag(tag ConstantPoolTag) bool {
	return tag == m.Tag
}

func (m ReferenceItem) write(w io.Writer, constantPool ConstantPool, _ int) error {
	err := m.Tag.write(w)
	if err != nil {
		return err
	}

	err = jvmio.WritePaddedBytes(w, constantPool.FindClassNameItem(m.Class), 2)
	if err != nil {
		return err
	}

	err = jvmio.WritePaddedBytes(w, constantPool.FindNameAndType(m.Name, m.Type), 2)
	if err != nil {
		return err
	}

	return nil
}

func isMethodReference(c constantPoolItem, class string, name string, typeDescriptor string) bool {
	return c.isTag(MethodReference) &&
		c.(ReferenceItem).Class == class &&
		c.(ReferenceItem).Name == name &&
		c.(ReferenceItem).Type == typeDescriptor
}

func isFieldReference(c constantPoolItem, class string, name string, typeDescriptor string) bool {
	return c.isTag(FieldReference) &&
		c.(ReferenceItem).Class == class &&
		c.(ReferenceItem).Name == name &&
		c.(ReferenceItem).Type == typeDescriptor
}
