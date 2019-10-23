package List

import "fmt"

type Node struct {
	Value int
	Next *Node
}

func Enter(l *Node, elem int) int {
	if l == nil {
		l = &Node{
			Value: elem,
			Next:  nil,
		}
		return 0
	}

	if elem == l.Value {
		fmt.Println("Value already in List:" , elem)
		return -1
	}

	if l.Next == nil {
		l.Next = &Node{elem, nil}
		return -2
	}
	return Enter(l.Next, elem)
}

func Traverse(t *Node) {
	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}
	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}