package compiler

import (
	"go-on-jvm/intermediate"
	"go-on-jvm/jvm"
	definitions "go-on-jvm/jvm/definitions"
)

type CompiledClass struct {
	Path  string
	Class definitions.Class
}

func Compile(parsed intermediate.Parsed) []CompiledClass {
	var compiledClasses []CompiledClass

	for _, p := range parsed.Packages {
		compiledClasses = append(compiledClasses, CompilePackage(p)...)
	}

	return compiledClasses
}

func CompilePackage(pkg intermediate.Package) []CompiledClass {
	var compiledClasses []CompiledClass

	for _, class := range pkg.Classes() {
		jvmClass := definitions.NewClass(class.Name, jvm.ObjectClass)

		if class.IsPublic() {
			jvmClass.WithAccess(definitions.Super, definitions.Public)
		} else {
			jvmClass.WithAccess(definitions.Super)
		}

		compileFields(&jvmClass, class.Fields)

		compiledClass := CompiledClass{
			Path:  jvmClass.Name,
			Class: jvmClass,
		}

		compiledClasses = append(compiledClasses, compiledClass)
	}

	return compiledClasses
}

func compileFields(class *definitions.Class, fields []intermediate.Field) {
	for _, field := range fields {

		var jvmType string

		switch field.Type {
		case "int":
			jvmType = jvm.Int

		default:
			jvmType = jvm.ObjectClass
		}

		jvmField := definitions.NewField(field.Name, jvmType)

		if field.IsPublic() {
			jvmField.WithAccess(definitions.Super, definitions.Public)
		} else {
			jvmField.WithAccess(definitions.Super)
		}

		class.AddField(jvmField)
	}
}
