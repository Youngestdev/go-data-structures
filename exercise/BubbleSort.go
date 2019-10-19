package main

func BubbleSort(a []int) []int {
	// First we declare `j` which is equal to the index of the top element.
	// Then loop j downwards if j > 0
	for j := len(a) - 1; 0 < j; j-- {
		// Initialize an int i, check it it is greater than j then increment i.e loop onwards.
		for i := 0; i < j; i++ {
			// if the index of the next value in the array is less than the previous index a[i], swap.
			// So basically the comparison done there is ermm.... the second element in the array is compared to first element which is the base index until the base index's value is the next element in the array, i.e a[i+1] in a continuous loop until the highest value is attained.
			//fmt.Println(a[i+1], a[i])
			if a[i+1] < a[i] {
				//	Swap values!
				a[i], a[i+1] = a[i+1], a[i]
				// Next we check if the list is ordered after a pass ( or swap )
				if a[i] > a[i+1] {
					break
				}
				continue
			}
		}
	}
	return a
}
