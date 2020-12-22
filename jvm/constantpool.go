package jvm

import (
	"io"
)

type constantTag byte

const (
	utf8ConstantTag        constantTag = 0x01
	intConstantTag         constantTag = 0x03
	floatConstantTag       constantTag = 0x04
	longConstantTag        constantTag = 0x05
	doubleConstantTag      constantTag = 0x06
	classConstantTag       constantTag = 0x07
	stringConstantTag      constantTag = 0x08
	fieldConstantTag       constantTag = 0x09
	methodConstantTag      constantTag = 0x0A
	interfaceConstantTag   constantTag = 0x0B
	nameAndTypeConstantTag constantTag = 0x0C
)

func (c constantTag) write(w io.Writer) error {
	_, err := w.Write([]byte{byte(c)})
	return err
}

type constantPoolItem interface {
	hasTag(tag constantTag) bool
	write(w io.Writer, index int) error
}

// Pool

type constantPool struct {
	Items []constantPoolItem
}

func (c constantPool) write(w io.Writer) error {
	for i, item := range c.Items {
		err := item.write(w, i)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *constantPool) findClassNameItem(name string) int {
	isClassName := false

	for i, item := range c.Items {
		if isClassName && isUtf8(item, name) {
			// We want the index of the class specifier, rather than the UTF8
			return i - 1
		}

		// If we're a class constant, then the next constant should be the class name
		isClassName = item.hasTag(classConstantTag)
	}

	return 0
}

func (c *constantPool) addClassItem(name string) {
	// No need to duplicate
	if c.findClassNameItem(name) > 0 {
		return
	}

	c.addItem(newClassConstant())
	c.addItem(newUtf8Constant(name))
}

func (c *constantPool) addItem(item constantPoolItem) {
	c.Items = append(c.Items, item)
}

func newConstantPool() *constantPool {
	c := &constantPool{
		make([]constantPoolItem, 0),
	}
	c.addItem(newPoolSize(c))
	return c
}

// Pool Size
type poolSize struct {
	constantPool *constantPool
}

func (p *poolSize) hasTag(_ constantTag) bool {
	return false
}

func (p *poolSize) write(w io.Writer, _ int) error {
	return writePaddedBytes(w, len(p.constantPool.Items), 2)
}

func newPoolSize(c *constantPool) *poolSize {
	return &poolSize{c}
}

// UTF8
// TODO Encode in JVM version of UTF8, rather than Unicode Standard

type utf8ConstantItem struct {
	Value string
}

func (u utf8ConstantItem) hasTag(tag constantTag) bool {
	return tag == utf8ConstantTag
}

func (u utf8ConstantItem) write(w io.Writer, _ int) error {
	err := utf8ConstantTag.write(w)
	if err != nil {
		return err
	}

	// We want the length in bytes, not Rune length
	err = writePaddedBytes(w, len(u.Value), 2)
	if err != nil {
		return err
	}
	_, err = io.WriteString(w, u.Value)
	return err
}

func newUtf8Constant(value string) utf8ConstantItem {
	return utf8ConstantItem{value}
}

func isUtf8(c constantPoolItem, value string) bool {
	return c.hasTag(utf8ConstantTag) && c.(utf8ConstantItem).Value == value
}

// Class

type classConstantItem struct{}

func (c classConstantItem) hasTag(tag constantTag) bool {
	return tag == classConstantTag
}

func (c classConstantItem) write(w io.Writer, index int) error {
	err := classConstantTag.write(w)
	if err != nil {
		return err
	}

	err = writePaddedBytes(w, index+1, 2)
	if err != nil {
		return err
	}

	return nil
}

func newClassConstant() classConstantItem {
	return classConstantItem{}
}
