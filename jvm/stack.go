package jvm

import "go-on-jvm/jvm/constantpool"

const codeAttribute = "Code"

type Stack struct {
	Arguments  []Variable
	Statements []Statement
}

func (s Stack) Empty() bool {
	return len(s.Statements) == 0
}

func (s Stack) Index(variable Variable) int {
	for i, v := range s.Variables() {
		if v == variable {
			return i
		}
	}
	return -1
}

func (s Stack) Variables() []Variable {
	variables := make([]Variable, len(s.Arguments))
	m := make(map[Variable]bool)

	for i, argument := range s.Arguments {
		m[argument] = true
		variables[i] = argument
	}

	var index int

	for _, statement := range s.Statements {
		for _, v := range statement.Variables() {
			if m[v] {
				continue
			}

			index++
			m[v] = true
			variables = append(variables, v)
		}
	}

	return variables
}

func (s Stack) MaxSize() int {
	max := len(s.Arguments)

	for _, statement := range s.Statements {
		stackSize := len(statement.Variables())

		if stackSize > max {
			max = stackSize
		}
	}

	return max
}

func (s Stack) MaxLocals() int {
	locals := 0

	for _, variable := range s.Variables() {
		if variable.IsLocal() {
			locals++
		}
	}

	return 0
}

func (s Stack) fillConstantsPool(pool *constantpool.ConstantPool) {
	if s.Empty() {
		return
	}

	pool.AddUTF8(codeAttribute)

	for _, statement := range s.Statements {
		statement.fillConstantsPool(pool)
	}
}
