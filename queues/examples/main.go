package main

import (
	"fmt"
	Queue "github.com/Youngestdev/go-data-structures/queues"
)

func main() {
	a := Queue.LinkedQueue{}
	a.Enter("Abdul")
	a.Enter("Adeshina")
	a.Leave()
	fmt.Println(a.Front())
	a.Leave()
	fmt.Println(a.Size())
	fmt.Println(a.Empty())
}