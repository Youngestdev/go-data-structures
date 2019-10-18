package Queue

import "errors"

type node struct {
	item interface{}
	next *node
}

type LinkedQueue struct {
	frontPtr, rearPtr *node
	count int
}

func (s *LinkedQueue) Enter(e interface{}) {
	s.frontPtr = &node{
		item: e,
		next: s.frontPtr,
	}
	s.rearPtr = &node{
		item: s.frontPtr.item,
		next: s.rearPtr,
	}
	s.count++
}

func (s *LinkedQueue) Leave() (interface{}, error) {
	if s.count == 0 {
		return nil, errors.New("leave: queue should'nt be empty")
	}

	result := s.frontPtr.item
	s.frontPtr = s.frontPtr.next
	s.count--
	return result, nil
}

func (s *LinkedQueue) Front() (interface{}, error) {
	if s.count == 0 {
		return nil, errors.New("top: error sha")
	}
	return s.frontPtr.item, nil
}

func (s *LinkedQueue) Empty() bool {
	return s.count == 0
}

func (s *LinkedQueue) Size() int {
	return s.count
}