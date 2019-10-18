package Stack

import "errors"

type node struct {
	item interface{}
	next *node
}

type LinkedStack struct {
	topPtr *node // singly-linked list of values
	count  int   // how many elements are present
}

//Size returns the number of elements in the LinkedStack
func (s *LinkedStack) Size() int {
	return s.count
}

// Empty() returns a true or false value
func (s *LinkedStack) Empty() bool {
	return s.count == 0
}

// Clear() removes/deletes every element in the stack
func (s *LinkedStack) Clear() {
	s.count = 0
	s.topPtr = nil
}

// Push() adds a new element into the stack
func (s *LinkedStack) Push(e interface{}) {
	// New nodes are added when new items are pushed into the stack,
	// hence, the nodes are updated directly.
	s.topPtr = &node{
		item: e,
		next: s.topPtr,
	}
	// increase the count of the items in the stack
	s.count++
}

// Pop() removes the last element in the stack
func (s *LinkedStack) Pop() (interface{}, error) {
	if s.count == 0 {
		return nil, errors.New("pop: the stack cannot be empty")
	}
	// remove the first item
	result := s.topPtr.item
	// update the second item as the first one
	s.topPtr = s.topPtr.next
	// update count
	s.count++
	return result, nil
}

// Top() returns the first item in the stack
func (s *LinkedStack) Top() (interface{}, error) {
	if s.count == 0 {
		return nil, errors.New("top: the stack cannot be empty")
	}
	return s.topPtr.item, nil
}
