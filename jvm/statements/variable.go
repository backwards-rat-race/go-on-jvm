package statements

import (
	"go-on-jvm/jvm/constantpool"
	jvmio "go-on-jvm/jvm/io"
	"go-on-jvm/jvm/opcodes"
	jvmtypes "go-on-jvm/jvm/types"
)

var (
	SelfReferenceVariable    = Variable{}
	GetSelfReferenceVariable = NewVariableGet(SelfReferenceVariable)
)

type Variable struct {
	Name string
	Type jvmtypes.TypeReference
}

func NewLocalVariable(name string, typeDescriptor jvmtypes.TypeReference) Variable {
	return Variable{Name: name, Type: typeDescriptor}
}

type VariableGet struct {
	Variable Variable
}

func (v VariableGet) GetInstructions(stack *Stack, pool *constantpool.ConstantPool) []byte {
	var instructions []byte

	index := stack.Load(v.Variable)
	instructions = opcodes.GetALoadInstruction(index)

	return instructions
}

func (v VariableGet) FillConstantsPool(_ *constantpool.ConstantPool) {
	// No constants
}

func (v VariableGet) MaxStack() uint {
	return 1
}

func NewVariableGet(variable Variable) VariableGet {
	return VariableGet{variable}
}

type VariableSet struct {
	Variable Variable
	Value    Statement
}

func (v VariableSet) GetInstructions(stack *Stack, pool *constantpool.ConstantPool) []byte {
	instructions := v.Value.GetInstructions(stack, pool)
	index := stack.Store(v.Variable)
	instructions = append(instructions, opcodes.GetAStoreInstruction(index)...)
	return instructions
}

func (v VariableSet) FillConstantsPool(pool *constantpool.ConstantPool) {
	v.Value.FillConstantsPool(pool)
}

func (v VariableSet) MaxStack() uint {
	if v.Value == nil {
		return 0
	} else {
		return v.Value.MaxStack()
	}
}

func NewVariableSet(variable Variable, value Statement) VariableSet {
	return VariableSet{variable, value}
}

type StaticVariable struct {
	Variable
	Class string
}

func NewStaticVariable(class string, name string, typeDescriptor jvmtypes.TypeReference) StaticVariable {
	return StaticVariable{
		Variable: NewLocalVariable(name, typeDescriptor),
		Class:    class,
	}
}

type StaticVariableGet struct {
	Variable StaticVariable
}

func (s StaticVariableGet) GetInstructions(_ *Stack, pool *constantpool.ConstantPool) []byte {
	instructions := []byte{opcodes.GETSTATIC}
	instructions = jvmio.AppendPaddedBytes(
		instructions,
		pool.FindFieldReference(
			s.Variable.Class,
			s.Variable.Name,
			s.Variable.Type.JvmRef(),
		),
		2,
	)
	return instructions
}

func (s StaticVariableGet) FillConstantsPool(pool *constantpool.ConstantPool) {
	pool.AddFieldReference(
		s.Variable.Class,
		s.Variable.Name,
		s.Variable.Type.JvmRef(),
	)
}

func (s StaticVariableGet) MaxStack() uint {
	return 1
}

func NewStaticVariableGet(variable StaticVariable) StaticVariableGet {
	return StaticVariableGet{variable}
}
