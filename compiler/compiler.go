package compiler

import (
	"go-on-jvm/intermediate"
	"go-on-jvm/jvm"
)

type CompiledClass struct {
	Path  string
	Class jvm.Class
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
		jvmClass := jvm.NewClass(class.Name, jvm.ObjectClass)

		if class.IsPublic() {
			jvmClass.WithAccess(jvm.Super, jvm.Public)
		} else {
			jvmClass.WithAccess(jvm.Super)
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

func compileFields(class *jvm.Class, fields []intermediate.Field) {
	for _, field := range fields {

		var jvmType string

		switch field.Type {
		case "int":
			jvmType = jvm.Int

		default:
			jvmType = jvm.ObjectClass
		}

		jvmField := jvm.NewField(field.Name, jvmType)

		if field.IsPublic() {
			jvmField.WithAccess(jvm.Super, jvm.Public)
		} else {
			jvmField.WithAccess(jvm.Super)
		}

		class.AddField(jvmField)
	}
}
