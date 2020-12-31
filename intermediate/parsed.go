package intermediate

import "fmt"

type Parsed struct {
	Packages []Package
}

func (p *Parsed) AddPackage(pkg Package) {
	p.Packages = append(p.Packages, pkg)
}

type Package struct {
	Name           string
	Encapsulations []Encapsulated
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

type Class struct {
	Fields []Field
}

func (c *Class) AddFields(fields []Field) {
	c.Fields = append(c.Fields, fields...)
}

func (c *Class) AddField(field Field) {
	c.Fields = append(c.Fields, field)
}
