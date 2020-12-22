package jvm

import (
	"go-on-jvm/jvm/constants"
	jvmio "go-on-jvm/jvm/io"
	"io"
)

type Class struct {
	Name       string
	Access     []AccessModifier
	Super      string
	Implements []string
	Methods    []Method
}

func NewClass(name string, super string) *Class {
	return &Class{Name: name, Super: super}
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

	constantPool := constants.NewConstantPool()
	c.fillConstantPool(constantPool)
	err = constantPool.Write(w)
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

func (c Class) fillConstantPool(pool *constants.ConstantPool) {
	pool.AddClassReference(c.Name)
	pool.AddClassReference(c.Super)
}

func (c Class) writeAccess(w io.Writer) error {
	var access int

	for _, modifier := range c.Access {
		access |= int(modifier)
	}

	return jvmio.WritePaddedBytes(w, access, 2)
}

func (c Class) writeClassSpecifier(w io.Writer, pool *constants.ConstantPool) error {
	err := jvmio.WritePaddedBytes(w, pool.FindClassNameItem(c.Name), 2)
	if err != nil {
		return err
	}

	err = jvmio.WritePaddedBytes(w, pool.FindClassNameItem(c.Super), 2)
	if err != nil {
		return err
	}

	_, err = w.Write([]byte{
		0x00, 0x00, // Interfaces count
		0x00, 0x00, // Fields count
	})

	err = jvmio.WritePaddedBytes(w, len(c.Methods), 2)
	if err != nil {
		return err
	}

	// Attributes count
	_, err = w.Write([]byte{0x00, 0x00})
	return err
}

func (c Class) writeMethods(w io.Writer, pool *constants.ConstantPool) error {
	for _, method := range c.Methods {
		method.fillConstantsPool(pool)

		err := method.Compile(w)
		if err != nil {
			return err
		}
	}
	return nil
}
