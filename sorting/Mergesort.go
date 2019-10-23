package sorting

func MergeSort(a []int) []int {
	aux := make([]int, len(a))
	copy(aux, a)
	// Recursive function to fulfil the `divide and conquer` tag.
	mergeInto(a, aux)
	// Return sorted list.
	return a
}

// MergeInto() takes two integer slices, divides them and then sort them before merging them back.

func mergeInto(dst []int, src []int ) {
	if len(dst) < 2 {
		return
	}
	m := len(dst) / 2

	// Recursive function to split the inputs and then merge them back after sorting..

	// Merge the the first halve of the slice m
	mergeInto(src[:m], dst[:m])
	// Merge the second halve of the slice m
	mergeInto(src[m:], dst[m:])

	j, k := 0, m

	for i := 0; i < len(dst); i++ {
		if j < m && k < len(src) {
			if src[j] < src[k] {
				dst[i], j = src[j], j+1
			} else {
				dst[i], k = src[k], k + 1
			}
		} else if j < m {
			dst[i], j = src[j], j + 1
		} else {
			dst[i], k = src[k], k + 1
		}
	}
}