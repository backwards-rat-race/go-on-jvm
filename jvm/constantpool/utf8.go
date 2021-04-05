package constantpool

import (
	jvmio "go-on-jvm/jvm/io"
	"io"
)

// TODO Encode in JVM version of UTF8, rather than Unicode Standard

type Utf8Item struct {
	Value string
}

func (u Utf8Item) isTag(tag ConstantPoolTag) bool {
	return tag == Utf8
}

func (u Utf8Item) write(w io.Writer, _ ConstantPool, _ int) error {
	err := Utf8.write(w)
	if err != nil {
		return err
	}

	// We want the length in bytes, not Rune length
	err = jvmio.WritePaddedBytesI(w, len(u.Value), 2)
	if err != nil {
		return err
	}
	_, err = io.WriteString(w, u.Value)
	return err
}

func newUtf8Constant(value string) Utf8Item {
	return Utf8Item{value}
}

func isUtf8(c constantPoolItem, value string) bool {
	return c.isTag(Utf8) && c.(Utf8Item).Value == value
}
