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
	// FIXME Should the int be a separate entry in the constant pool? See String Constant
	err := Int.write(w)
	if err != nil {
		return err
	}

	err = jvmio.WritePaddedBytesI(w, i.Constant, 1)
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

type StringConstantItem struct {
	Constant string
}

func (s StringConstantItem) isTag(tag ConstantPoolTag) bool {
	return tag == StringReference
}

func (s StringConstantItem) write(w io.Writer, constantPool ConstantPool, index int) error {
	err := StringReference.write(w)
	if err != nil {
		return err
	}

	err = jvmio.WritePaddedBytes(w, constantPool.FindUTF8Item(s.Constant), 2)
	if err != nil {
		return err
	}

	return nil
}

func newStringConstant(constant string) StringConstantItem {
	return StringConstantItem{Constant: constant}
}

func isString(c constantPoolItem, constant string) bool {
	return c.isTag(StringReference) && c.(StringConstantItem).Constant == constant
}
