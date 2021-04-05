package statements

type Stack struct {
	Locals []Variable
}

func NewStack(arguments ...Variable) *Stack {
	return &Stack{arguments}
}

func (s Stack) Load(variable Variable) int {
	for i, v := range s.Locals {
		if v == variable {
			return i
		}
	}
	return -1
}

func (s *Stack) Store(variable Variable) int {
	index := s.Load(variable)
	if index > -1 {
		return index
	}
	index = len(s.Locals)
	s.Locals = append(s.Locals, variable)
	return index
}

func (s Stack) MaxLocals() int {
	return len(s.Locals)
}
