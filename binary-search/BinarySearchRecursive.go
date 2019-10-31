package binary_search

func BinarySearchRecursive(a []int, key int) (int, bool) {
	if len(a) == 0 { return -1, false}
	m := len(a) / 2
	switch {
	case key == a[m]:
		return m, true
	case key < a[m]: return BinarySearchRecursive(a[:m], key)
	case key > a[m]: return BinarySearchRecursive(a[m+1:], key)
	}
	panic("Unreachable code has been reached")
}
