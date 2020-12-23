package jvm

import (
	jvmio "go-on-jvm/jvm/io"
	"io"
	"strings"
)

type AccessModifier int

const (
	Public = 1 << iota
	_
	_
	_
	Final
	Super
	Interface
	Abstract
	Synthetic
	Annotation
	Enum
)

const (
	Void   = "V"
	Int    = "I"
	Float  = "F"
	Double = "D"

	ClassRef = "L"
)

const ObjectClass = "java/lang/Object"

type Instruction interface {
	Compile(w io.Writer) error
}

func JavaPackageToJvmPackage(javaPackageName string) string {
	return strings.Join(strings.Split(javaPackageName, "."), "/")
}

func writeAccessModifier(w io.Writer, accessModifiers []AccessModifier) error {
	var access int

	for _, modifier := range accessModifiers {
		access |= int(modifier)
	}

	return jvmio.WritePaddedBytes(w, access, 2)
}
