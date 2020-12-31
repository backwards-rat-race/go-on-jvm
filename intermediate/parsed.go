package intermediate

import (
	"fmt"
	"unicode"
)

func isPublicName(name string) bool {
	return len(name) > 0 && unicode.IsUpper(rune(name[0]))
}

type Parsed struct {
	Packages []Package
}

func (p *Parsed) AddPackage(pkg Package) {
	p.Packages = append(p.Packages, pkg)
}

func (p *Parsed) Classes() []Class {
	var classes []Class
	for _, pkg := range p.Packages {
		classes = append(classes, pkg.Classes()...)
	}
	return classes
}

type Package struct {
	Name           string
	Encapsulations []Encapsulated
}

func (p *Package) Classes() []Class {
	var classes []Class
	for _, encapsulation := range p.Encapsulations {
		classes = append(encapsulation.Classes, classes...)
	}
	return classes
}

func (p *Package) AddEncapsulation(encapsulated Encapsulated) error {
	if encapsulated.Package != p.Name {
		return fmt.Errorf("encapsulation package '%s' does not match parent package '%s'", encapsulated.Package, p.Name)
	}
	p.Encapsulations = append(p.Encapsulations, encapsulated)
	return nil
}

type Encapsulated struct {
	Package   string
	Imports   []Import
	Variables []VariableGroup
	Classes   []Class
}

func (d *Encapsulated) AddImport(declarationImport Import) {
	d.Imports = append(d.Imports, declarationImport)
}

func (d *Encapsulated) AddVariableGroup(variableGroup VariableGroup) {
	d.Variables = append(d.Variables, variableGroup)
}

func (d *Encapsulated) AddClass(class Class) {
	d.Classes = append(d.Classes, class)
}

type ParsedClass struct {
}

type Import struct {
	Alias   string
	Package string
}

type VariableGroup struct {
	Const     bool
	Variables []Variable
}

func (v *VariableGroup) AddVariable(variable Variable) {
	v.Variables = append(v.Variables, variable)
}

type Variable struct {
	Name  string
	Type  string
	Value string
}

type Field struct {
	Name     string
	Type     string
	TypeOnly bool
}

func (f Field) IsPublic() bool {
	return isPublicName(f.Name)
}

type Class struct {
	Name   string
	Fields []Field
}

func (c Class) IsPublic() bool {
	return isPublicName(c.Name)
}

func (c *Class) AddFields(fields []Field) {
	c.Fields = append(c.Fields, fields...)
}

func (c *Class) AddField(field Field) {
	c.Fields = append(c.Fields, field)
}
