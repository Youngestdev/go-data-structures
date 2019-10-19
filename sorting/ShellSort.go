package sorting

func ShellSort(n []int) []int {
	h := 1
	for h < len(n)/9 {
		h = 3 * h + 1
	}

	for 0 < h {
		for j := h; j < len(n); j++ {
			element := n[j]
			var i int
			for i = j; h <= i && element < n[i-h];  i -= h {
				n[i] = n[i-h]
			}
			n[i] = element
		}
		h /= 3
	}
	return n
}