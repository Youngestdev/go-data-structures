package sorting

func InsertionSort(a []int) []int {
	// We'll start from the second element.
	for j := 1; j < len(a); j++ {
		// Inserted elements are stored in `element`
		element := a[j]

		var i int
		// If i = 1, the first value won't be present.
		for i = j; 0 < i && element < a[i-1]; i-- {
			a[i]= a[i-1]
		}
		a[i] = element
	}
	return a
}