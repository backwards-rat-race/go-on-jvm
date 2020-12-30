package structure

type Parsed struct {
	Packages []Package
}

func (p *Parsed) AddPackage(pkg Package) {
	p.Packages = append(p.Packages, pkg)
}

type Package struct {
	Name                 string
	DeclarationsContexts []DeclarationContext
}

func (p *Package) AddDeclarationsContext(declarationsContext DeclarationContext) {
	p.DeclarationsContexts = append(p.DeclarationsContexts, declarationsContext)
}

type DeclarationContext struct {
	Package   string
	Imports   []Import
	Variables []VariableGroup
	Classes   []Class
}

func (d *DeclarationContext) AddImport(declarationImport Import) {
	d.Imports = append(d.Imports, declarationImport)
}

func (d *DeclarationContext) AddVariableGroup(variableGroup VariableGroup) {
	d.Variables = append(d.Variables, variableGroup)
}

func (d *DeclarationContext) AddClass(class Class) {
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
