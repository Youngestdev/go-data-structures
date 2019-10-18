package stacks_and_recursion

func RecursiveSearch(slice []int, value int) bool {
	if slice == nil || len(slice) == 0 {
		return false
	}
	if slice[0] == value {
		return true
	}
	return RecursiveSearch(slice[1:], value)
}

func Search(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
