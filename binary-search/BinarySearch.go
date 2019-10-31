package binary_search

func BinarySearch(a []int, key int) (int, bool) {
	for lb, ub := 0, len(a)-1; lb <= ub; {
		m := len(a)/2
		switch {
		case key == a[m]: return m, true
		case key < a[m]: ub = m - 1
		case key > a[m]: lb = m +1
		}
	}
	return -1, false
}