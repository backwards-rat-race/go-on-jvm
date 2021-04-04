package constantpool

import (
	jvmio "go-on-jvm/jvm/io"
	"io"
)

type IntConstantItem struct {
	Constant int
}

func (i IntConstantItem) isTag(tag ConstantPoolTag) bool {
	return tag == Int
}

func (i IntConstantItem) write(w io.Writer, _ ConstantPool, _ int) error {
	err := Int.write(w)
	if err != nil {
		return err
	}

	err = jvmio.WritePaddedBytes(w, i.Constant, 1)
	if err != nil {
		return err
	}

	return nil
}

func newIntConstant(constant int) IntConstantItem {
	return IntConstantItem{Constant: constant}
}

func isInt(c constantPoolItem, constant int) bool {
	return c.isTag(Int) && c.(IntConstantItem).Constant == constant
}
