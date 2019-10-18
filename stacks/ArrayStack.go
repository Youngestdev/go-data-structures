package Stack

import "errors"

type ArrayStack struct {
	store []interface{}
}

func (s *ArrayStack) Size() int {
	return len(s.store)
}

func (s *ArrayStack) Empty() bool {
	return len(s.store) == 0
}

func (s *ArrayStack) Clear() {
	s.store = make([]interface{}, 0, 10)
}

func (s *ArrayStack) Push(e interface{}) {
	s.store = append(s.store, e)
}

func (s *ArrayStack) Pop() (interface{}, error) {
	if len(s.store) == 0 {
		return nil, errors.New("pop: the stack cannot be empty")
	}
	res := s.store[len(s.store)-1]
	s.store = s.store[:len(s.store)-1]
	return res, nil
}

func (s *ArrayStack) Top() (interface{}, error) {
	if len(s.store) == 0 {
		return nil, errors.New("top: stack cannot be empty")
	}
	return s.store[:len(s.store)-1], nil
}
