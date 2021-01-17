package constantpool

import (
	jvmio "go-on-jvm/jvm/io"
	"io"
)

type ClassReferenceItem struct {
	Name string
}

func newClassConstant(name string) ClassReferenceItem {
	return ClassReferenceItem{name}
}

func (c ClassReferenceItem) isTag(tag ConstantPoolTag) bool {
	return tag == ClassReference
}

func (c ClassReferenceItem) write(w io.Writer, constantPool ConstantPool, _ int) error {
	err := ClassReference.write(w)
	if err != nil {
		return err
	}

	err = jvmio.WritePaddedBytes(w, constantPool.FindUTF8Item(c.Name), 2)
	if err != nil {
		return err
	}

	return nil
}

func isClass(c constantPoolItem, name string) bool {
	return c.isTag(ClassReference) && c.(ClassReferenceItem).Name == name
}
