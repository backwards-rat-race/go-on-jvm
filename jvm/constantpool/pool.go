package constantpool

import (
	jvmio "go-on-jvm/jvm/io"
	"io"
)

type ConstantPoolTag int

const (
	_ ConstantPoolTag = iota
	Utf8
	_
	Int
	Float
	Long
	Double
	ClassReference
	StringReference
	FieldReference
	MethodReference
	InterfaceMethodReference
	NameAndTypeDescriptor
	_
	_
	MethodHandle
	MethodType
	Dynamic
	InvokeDynamic
	Module
	Package
)

func (c ConstantPoolTag) write(w io.Writer) error {
	_, err := w.Write([]byte{byte(c)})
	return err
}

type constantPoolItem interface {
	isTag(tag ConstantPoolTag) bool
	write(w io.Writer, constantPool ConstantPool, index int) error
}

// Pool

type ConstantPool struct {
	Items []constantPoolItem
}

func (c ConstantPool) Write(w io.Writer) error {
	for i, item := range c.Items {
		err := item.write(w, c, i)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *ConstantPool) FindClassNameItem(name string) int {
	return c.findItem(func(item constantPoolItem) bool {
		return isClass(item, name)
	})
}

func (c *ConstantPool) FindUTF8Item(value string) int {
	return c.findItem(func(item constantPoolItem) bool {
		return isUtf8(item, value)
	})
}

func (c *ConstantPool) FindMethodReference(class string, name string, typeDescriptor string) int {
	return c.findItem(func(item constantPoolItem) bool {
		return isMethodReference(item, class, name, typeDescriptor)
	})
}

func (c *ConstantPool) FindNameAndType(name string, typeDescriptor string) int {
	return c.findItem(func(item constantPoolItem) bool {
		return isNameAndType(item, name, typeDescriptor)
	})
}

func (c *ConstantPool) FindIntConstant(constant int) int {
	return c.findItem(func(item constantPoolItem) bool {
		return isInt(item, constant)
	})
}

func (c *ConstantPool) findItem(predicate func(item constantPoolItem) bool) int {
	for i, item := range c.Items {
		if predicate(item) {
			return i
		}
	}

	return 0
}

func (c *ConstantPool) AddClassReference(name string) {
	if c.FindClassNameItem(name) > 0 {
		return
	}

	c.addItem(newClassConstant(name))
	c.AddUTF8(name)
}

func (c *ConstantPool) AddUTF8(value string) {
	if c.FindUTF8Item(value) > 0 {
		return
	}
	c.addItem(newUtf8Constant(value))
}

func (c *ConstantPool) AddMethodReference(class string, name string, typeDescriptor string) {
	if c.FindMethodReference(class, name, typeDescriptor) > 0 {
		return
	}
	c.addItem(newMethodReference(class, name, typeDescriptor))
	c.AddClassReference(class)
	c.AddNameAndType(name, typeDescriptor)
}

func (c *ConstantPool) AddNameAndType(name string, typeDescriptor string) {
	if c.FindNameAndType(name, typeDescriptor) > 0 {
		return
	}
	c.addItem(newNameAndType(name, typeDescriptor))
	c.AddUTF8(name)
	c.AddUTF8(typeDescriptor)
}

func (c *ConstantPool) AddIntConstant(constant int) {
	if c.FindIntConstant(constant) > 0 {
		return
	}
	c.addItem(newIntConstant(constant))
}

func (c *ConstantPool) addItem(item constantPoolItem) {
	c.Items = append(c.Items, item)
}

func New() *ConstantPool {
	c := &ConstantPool{}
	c.addItem(poolSize{})
	return c
}

// Pool Size
type poolSize struct{}

func (p poolSize) isTag(_ ConstantPoolTag) bool {
	return false
}

func (p poolSize) write(w io.Writer, constantPool ConstantPool, _ int) error {
	return jvmio.WritePaddedBytes(w, len(constantPool.Items), 2)
}
