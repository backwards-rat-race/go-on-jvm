package definitions

import (
	jvmio "go-on-jvm/jvm/io"
	"io"
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

func writeAccessModifier(w io.Writer, accessModifiers []AccessModifier) error {
	var access uint

	for _, modifier := range accessModifiers {
		access |= uint(modifier)
	}

	return jvmio.WritePaddedBytes(w, access, 2)
}
