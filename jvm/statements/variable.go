package statements

import (
	"go-on-jvm/jvm/constantpool"
	"go-on-jvm/jvm/opcodes"
	jvmtypes "go-on-jvm/jvm/types"
)

var (
	SelfReferenceVariable    = Variable{}
	GetSelfReferenceVariable = NewVariableGet(SelfReferenceVariable)
)

type Variable struct {
	Object string
	Name   string
	Type   jvmtypes.TypeReference
}

func NewLocalVariable(name string, typeDescriptor jvmtypes.TypeReference) Variable {
	return Variable{Name: name, Type: typeDescriptor}
}

func NewObjectVariable(object string, name string, typeDescriptor jvmtypes.TypeReference) Variable {
	return Variable{object, name, typeDescriptor}
}

func (v Variable) IsLocal() bool {
	return v.Object == ""
}

type VariableGet struct {
	Variable Variable
}

func (v VariableGet) GetInstructions(_ int, stack *Stack, _ *constantpool.ConstantPool) []byte {
	var instructions []byte

	if v.Variable.IsLocal() {
		index := stack.Load(v.Variable)
		instructions = opcodes.GetALoadInstruction(index)
	} else {
		// TODO
		panic("implement me")
	}
	stack.Push()
	return instructions
}

func (v VariableGet) FillConstantsPool(_ *constantpool.ConstantPool) {
	// No constants
}

func NewVariableGet(variable Variable) VariableGet {
	return VariableGet{variable}
}

type VariableSet struct {
	Variable Variable
	Value    Statement
}

func (v VariableSet) GetInstructions(writeIndex int, stack *Stack, pool *constantpool.ConstantPool) []byte {
	instructions := v.Value.GetInstructions(writeIndex, stack, pool)
	index := stack.Store(v.Variable)
	instructions = append(instructions, opcodes.GetAStoreInstruction(index)...)
	stack.Pop()
	return instructions
}

func (v VariableSet) FillConstantsPool(pool *constantpool.ConstantPool) {
	v.Value.FillConstantsPool(pool)
}

func NewVariableSet(variable Variable, value Statement) VariableSet {
	return VariableSet{variable, value}
}
