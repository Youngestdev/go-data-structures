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
	fmt.Println(" =============== Merge Sort ========== ")
	d := sorting.MergeSort([]int{1, 2, 4768, 56, 23, 56,12, 45, 6, 8, 9 ,0, 7,2,3,4,1})
	fmt.Println(d)
	fmt.Println(" =============== Quick Sort ========== ")
	e := sorting.QuickSort([]int{1, 2, 4768, 56, 23, 56,12, 45, 6, 8, 9 ,0, 7,2,3,4,1})
	fmt.Println(e)
}