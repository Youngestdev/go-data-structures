package main

import (
	"fmt"
	Queue "github.com/Youngestdev/go-data-structures/queues"
)

func main() {
	fmt.Println("================ LinkedQueue 1 =======================")
	a := Queue.LinkedQueue{}
	a.Enter("Abdul")
	a.Enter("Adeshina")
	_, _ = a.Leave()
	fmt.Println(a.Front())
	_, _ = a.Leave()
	fmt.Println(a.Size())
	fmt.Println(a.Empty())
	fmt.Println("================ LinkedQueue 2 =======================")
	b := Queue.ArrayQueue{}
	b.Enter("First things first")
	b.Enter("Second thing second hm")
	fmt.Println(b.Leave())
	fmt.Println(b.Top())
	fmt.Println(b.Size())
}
