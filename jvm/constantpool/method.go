package constantpool

import (
	jvmio "go-on-jvm/jvm/io"
	"io"
)

type MethodReferenceItem struct {
	Class string
	Name  string
	Type  string
}

func newMethodReference(class string, name string, typeDescriptor string) MethodReferenceItem {
	return MethodReferenceItem{class, name, typeDescriptor}
}

func (m MethodReferenceItem) isTag(tag ConstantPoolTag) bool {
	return tag == MethodReference
}

func (m MethodReferenceItem) write(w io.Writer, constantPool ConstantPool, _ int) error {
	err := MethodReference.write(w)
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
		c.(MethodReferenceItem).Class == class &&
		c.(MethodReferenceItem).Name == name &&
		c.(MethodReferenceItem).Type == typeDescriptor
}
