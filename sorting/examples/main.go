package main

import (
	"fmt"
	"github.com/Youngestdev/go-data-structures/sorting"
)

func main() {
	fmt.Println(" ==== Bubble Sort ==== ")
	a := sorting.BubbleSort([]int{7,2,3,4,1})
	fmt.Println(a, "\t")
	fmt.Println(" ==== Sorting Sort ==== ")
	b := sorting.SelectionSort([]int{7,2,3,4,1})
	fmt.Println(b)
	fmt.Println("============ Insertion Sort ========= ")
	c := sorting.InsertionSort([]int{7,2,3,4,1})
	fmt.Println(c)
}