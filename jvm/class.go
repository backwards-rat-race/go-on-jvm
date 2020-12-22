package jvm

import (
	"io"
)

type Class struct {
	Name       string
	Access     []AccessModifier
	Super      string
	Implements []string
	Methods    []Method
}

func (c *Class) WithAccess(modifier ...AccessModifier) {
	c.Access = modifier
}

func (c Class) Compile(w io.Writer) (err error) {
	err = c.writeMagic(w)
	if err != nil {
		return
	}

	err = c.writeVersion(w)
	if err != nil {
		return
	}

	constantPool := newConstantPool()
	c.fillConstantPool(constantPool)
	err = constantPool.write(w)
	if err != nil {
		return
	}

	err = c.writeAccess(w)
	if err != nil {
		return
	}

	err = c.writeClassSpecifier(w, constantPool)
	if err != nil {
		return
	}

	return nil
}

func (c Class) writeMagic(w io.Writer) error {
	_, err := w.Write([]byte{0xCA, 0xFE, 0xBA, 0xBE})
	return err
}

func (c Class) writeVersion(w io.Writer) error {
	// Java Version 8
	_, err := w.Write([]byte{0x00, 0x00, 0x00, 0x34})
	return err
}

func (c Class) fillConstantPool(pool *constantPool) {
	pool.addClassItem(c.Name)
	pool.addClassItem(c.Super)
}

func (c Class) writeAccess(w io.Writer) error {
	var access int

	for _, modifier := range c.Access {
		access |= int(modifier)
	}

	return writePaddedBytes(w, access, 2)
}

func (c Class) writeClassSpecifier(w io.Writer, pool *constantPool) error {
	err := writePaddedBytes(w, pool.findClassNameItem(c.Name), 2)
	if err != nil {
		return err
	}

	err = writePaddedBytes(w, pool.findClassNameItem(c.Super), 2)
	if err != nil {
		return err
	}

	_, err = w.Write([]byte{
		0x00, 0x00, // Interfaces count
		0x00, 0x00, // Fields count
		0x00, 0x00, // Methods count
		0x00, 0x00, // Attributes count
	})
	return err
}

func NewClass(name string, super string) *Class {
	return &Class{Name: name, Super: super}
}
