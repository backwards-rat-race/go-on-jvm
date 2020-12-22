package jvm

import (
	"io"
	"strings"
)

type AccessModifier int

const (
	Public     AccessModifier = 0x001
	Final      AccessModifier = 0x0010
	Super      AccessModifier = 0x0020
	Interface  AccessModifier = 0x0200
	Abstract   AccessModifier = 0x0400
	Synthetic  AccessModifier = 0x1000
	Annotation AccessModifier = 0x2000
	Enum       AccessModifier = 0x4000
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

func toPaddedBytes(seq int, bytes int) []byte {
	buf := make([]byte, bytes)
	for i := len(buf) - 1; seq != 0; i-- {
		buf[i] = byte(seq & 0xff)
		seq >>= 8
	}
	return buf
}

func writePaddedBytes(w io.Writer, seq int, bytes int) error {
	_, err := w.Write(toPaddedBytes(seq, bytes))
	return err
}
