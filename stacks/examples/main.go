package main

import (
	"fmt"
	Stack "github.com/Youngestdev/go-data-structures/stacks"
)

func main() {
	stack := Stack.LinkedStack{}
	stack.Push("Abdulazeez")
	stack.Push("Adeshina")
	stack.Push("Azeez")
	fmt.Println(stack.Size())
	fmt.Println(stack.Pop())
	stack.Clear()
	fmt.Println(stack.Top())
	fmt.Println("===========================Go=========== \t")
	stack2 := Stack.ArrayStack{}
	stack2.Push([]int{1, 2, 3, 4, 5, 6, 7})
	stack2.Push("stack")
	fmt.Println(stack2.Top())
}
