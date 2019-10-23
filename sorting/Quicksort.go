package sorting

func QuickSort(a []int) []int {
	// If this has less than 2 integers, return nothing.

	if len(a) < 2 {
		return nil
	}

	// We create a variable to hold `len(a)-1`

	ub := len(a) - 1

	// Pivot variable equals the value at index of ub in the array `a`.

	pivot := a[ub]

	// Declare lop variables.

	i, j := -1, ub

	// Loop

	for i < j {
		// Do nothing for i && j.
		for i++; a[i] < pivot; i++ {}
		for j--; 0 < j && a[j] > pivot; j-- {}

		// Swap values.

		a[i], a[j] = a[j], a[i]
	}
	// Swap values

	a[j], a[i], a[ub] = a[i], pivot, a[j]

	// Continue this - recursively.
	QuickSort(a[:i])
	QuickSort(a[i+1:])

	return a
}