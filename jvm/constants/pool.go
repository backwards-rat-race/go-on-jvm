package constants

import (
	jvmio "go-on-jvm/jvm/io"
	"io"
)

type PoolConstantTag int

const (
	_ PoolConstantTag = iota
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

func (c PoolConstantTag) write(w io.Writer) error {
	_, err := w.Write([]byte{byte(c)})
	return err
}

type constantPoolItem interface {
	isTag(tag PoolConstantTag) bool
	write(w io.Writer, index int) error
}

// Pool

type ConstantPool struct {
	Items []constantPoolItem
}

func (c ConstantPool) Write(w io.Writer) error {
	for i, item := range c.Items {
		err := item.write(w, i)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *ConstantPool) FindClassNameItem(name string) int {
	isClassName := false

	for i, item := range c.Items {
		if isClassName && isUtf8(item, name) {
			// We want the index of the class specifier, rather than the UTF8
			return i - 1
		}

		// If we're a class constant, then the next constant should be the class name
		isClassName = item.isTag(ClassReference)
	}

	return 0
}

func (c *ConstantPool) AddClassReference(name string) {
	// No need to duplicate
	if c.FindClassNameItem(name) > 0 {
		return
	}

	c.addItem(newClassConstant())
	c.addItem(newUtf8Constant(name))
}

func (c *ConstantPool) addItem(item constantPoolItem) {
	c.Items = append(c.Items, item)
}

func NewConstantPool() *ConstantPool {
	c := &ConstantPool{
		make([]constantPoolItem, 0),
	}
	c.addItem(newPoolSize(c))
	return c
}

// Pool Size
type poolSize struct {
	constantPool *ConstantPool
}

func (p *poolSize) isTag(_ PoolConstantTag) bool {
	return false
}

func (p *poolSize) write(w io.Writer, _ int) error {
	return jvmio.WritePaddedBytes(w, len(p.constantPool.Items), 2)
}

func newPoolSize(c *ConstantPool) *poolSize {
	return &poolSize{c}
}
