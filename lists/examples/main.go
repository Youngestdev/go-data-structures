package main

import (
	"fmt"
	List "github.com/Youngestdev/go-data-structures/lists"
)

func main() {
	var a List.ArrayList
	_, _ = a.Insert("Abdul")
	_, _ = a.Insert("Azeez")
	a.Put(0, "Adeshina")
	fmt.Println(a.Size())
	fmt.Println(a.Delete(0))
	fmt.Println(a.Index("Azeez"))
	_, _ = a.Insert("Azeez")
	fmt.Println(a.Size())
}
