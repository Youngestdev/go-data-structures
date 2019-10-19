package main

import "fmt"

func seqSearch(key int, ar []int) (bool, int) {
	for idx, k := range ar {
		if key == k {
			return true, idx
		}
	}
	return false, -1
}

func main()  {
	fmt.Println(seqSearch(16, []int{1,3,4,6,7}))
}