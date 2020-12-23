package constants

import (
	jvmio "go-on-jvm/jvm/io"
	"io"
)

type ClassReferenceItem struct{}

func newClassConstant() ClassReferenceItem {
	return ClassReferenceItem{}
}

func (c ClassReferenceItem) isTag(tag ConstantPoolTag) bool {
	return tag == ClassReference
}

func (c ClassReferenceItem) write(w io.Writer, index int) error {
	err := ClassReference.write(w)
	if err != nil {
		return err
	}

	err = jvmio.WritePaddedBytes(w, index+1, 2)
	if err != nil {
		return err
	}

	return nil
}
