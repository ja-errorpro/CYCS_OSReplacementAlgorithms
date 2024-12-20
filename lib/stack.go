package lib

type Stack struct {
	list []interface{}
}

func NewStack() *Stack {
	s := new(Stack)
	s.list = make([]interface{}, 0)
	return s
}

func (s *Stack) Push(v interface{}) {
	s.list = append(s.list, v)
}

func (s *Stack) Pop() interface{} {
	if len(s.list) == 0 {
		return nil
	}
	v := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]
	return v
}

func (s *Stack) Len() int {
	return len(s.list)
}

func (s *Stack) Empty() bool {
	return len(s.list) == 0
}

func (s *Stack) Top() interface{} {
	if len(s.list) == 0 {
		return nil
	}
	return s.list[len(s.list)-1]
}
