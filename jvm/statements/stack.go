package statements

type Stack struct {
	Locals      []Variable
	CurrentSize int

	maxSize int
}

func NewStack(arguments ...Variable) *Stack {
	return &Stack{arguments, 0, 0}
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

func (s *Stack) Pop() {
	s.CurrentSize -= 1
}

func (s *Stack) Push() {
	s.CurrentSize += 1
	if s.CurrentSize > s.maxSize {
		s.maxSize = s.CurrentSize
	}
}

func (s Stack) MaxSize() int {
	return s.maxSize
}

func (s Stack) MaxLocals() int {
	return len(s.Locals)
}
