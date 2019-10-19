package sorting

// Selection sort isn't quite different from Bubble sort except that it doesn't make many swaps during comparisons. This is shown on line 9 to 13
func SelectionSort(a []int) []int {
	for j := 0; j < len(a)-1; j++ {
		minIndex := j
		for i := j + 1; i < len(a); i++ {
			if a[i] < a[minIndex] {
				minIndex = i
			}
		}
		a[j], a[minIndex] = a[minIndex], a[j]
	}
	return a
}