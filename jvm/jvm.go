package jvm

import (
	"strings"
)

const (
	Void   = "V"
	Int    = "I"
	Float  = "F"
	Double = "D"

	ClassRef = "L"
)

const ObjectClass = "java/lang/Object"
const ConstructorName = "<init>"

func JavaPackageToJvmPackage(javaPackageName string) string {
	return strings.Join(strings.Split(javaPackageName, "."), "/")
}
