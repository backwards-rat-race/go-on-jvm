package types

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	dotReference = regexp.MustCompile(`^\[?([a-zA-Z_$][a-zA-Z\d_$]*\.*)*([a-zA-Z_$][a-zA-Z\d_$]*)$`)
	jvmClassRef  = regexp.MustCompile(`^L([a-zA-Z_$][a-zA-Z\d_$]*/*)*([a-zA-Z_$][a-zA-Z\d_$]*);$`)
	jvmArrayRef  = regexp.MustCompile(`^\[L([a-zA-Z_$][a-zA-Z\d_$]*/*)*([a-zA-Z_$][a-zA-Z\d_$]*);$`)
	jvmArray     = regexp.MustCompile(`^\[([a-zA-Z_$][a-zA-Z\d_$]*/*)*([a-zA-Z_$][a-zA-Z\d_$]*)$`)
	jvmReference = regexp.MustCompile(`^\[?([a-zA-Z_$][a-zA-Z\d_$]*/*)*([a-zA-Z_$][a-zA-Z\d_$]*)$`)
)

type TypeReference struct {
	reference string // JVM reference to an object. e.g. java/lang/Object (vs java.lang.Object)
}

func (t TypeReference) Dot() string {
	return strings.ReplaceAll(t.reference, "/", ".")
}

func (t TypeReference) Jvm() string {
	return t.reference
}

func (t TypeReference) JvmRef() string {
	if t.IsPrimitive() || t.IsArray() {
		return t.Jvm()
	} else {
		return classRef + t.Jvm() + ";"
	}
}

func (t TypeReference) Array() TypeReference {
	return MustParse(array + t.JvmRef())
}

func (t TypeReference) IsArray() bool {
	return strings.HasPrefix(t.Jvm(), array)
}

func (t TypeReference) IsPrimitive() bool {
	switch t.Jvm() {
	case int:
		fallthrough
	case float:
		fallthrough
	case double:
		fallthrough
	case void:
		return true
	default:
		return false
	}
}

func Parse(typeName string) (TypeReference, error) {
	switch {
	case jvmClassRef.MatchString(typeName):
		jvmReference := strings.TrimSuffix(strings.TrimPrefix(typeName, "L"), ";")
		return TypeReference{reference: jvmReference}, nil

	case jvmArray.MatchString(typeName):
		fallthrough
	case jvmArrayRef.MatchString(typeName):
		fallthrough
	case jvmReference.MatchString(typeName):
		return TypeReference{reference: typeName}, nil

	case dotReference.MatchString(typeName):
		jvmReference := strings.ReplaceAll(typeName, ".", "/")
		return TypeReference{reference: jvmReference}, nil

	default:
		return TypeReference{}, fmt.Errorf("invalid type reference: %s", typeName)
	}
}

func MustParse(typeName string) TypeReference {
	reference, err := Parse(typeName)
	if err != nil {
		panic(err)
	}
	return reference
}

type MethodType struct {
	Arguments  []TypeReference
	ReturnType TypeReference
}

func (m MethodType) Descriptor() string {
	typeDescriptor := "("

	for _, argType := range m.Arguments {
		typeDescriptor += argType.JvmRef()
	}

	typeDescriptor += ")"
	typeDescriptor += m.ReturnType.JvmRef()

	return typeDescriptor
}
