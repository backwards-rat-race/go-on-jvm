package jvm

import (
	jvmio "go-on-jvm/jvm/io"
	"io"
	"strings"
)

type AccessModifier int

const (
	Public AccessModifier = 1 << iota
	Private
	Protected
	Static
	Final
	Super
	Bridge
	Transient
	Native
	Interface
	Abstract
	Strict
	Synthetic
	Annotation
	Enum

	Synchronised = Super
	Volatile     = Bridge
	VarArgs      = Transient
)

const (
	Void   = "V"
	Int    = "I"
	Float  = "F"
	Double = "D"

	ClassRef = "L"
)

const ObjectClass = "java/lang/Object"

type Serialisable interface {
	Write(w io.Writer) error
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
