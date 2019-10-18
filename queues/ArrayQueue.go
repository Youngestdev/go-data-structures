package Queue

import "errors"

type ArrayQueue struct {
	queue []interface{}
}

func (s *ArrayQueue) Size() int {
	return len(s.queue)
}

func (s *ArrayQueue) Enter(e interface{}) {
	s.queue = append(s.queue, e)
}

func (s *ArrayQueue) Leave() (interface{}, error) {
	if len(s.queue) == 0 {
		return nil, errors.New("leave: queue shouldn't be empty")
	}
	res := s.queue[:len(s.queue)-1]
	// Not sure why this returns [] sha.
	s.queue = s.queue[:len(s.queue)-1]
	return res, nil
}

func (s *ArrayQueue) Top() (interface{}, error) {
	if len(s.queue) == 0 {
		return nil, errors.New("top: queue can't be empty")
	}
	return s.queue[:len(s.queue)-1], nil
}
