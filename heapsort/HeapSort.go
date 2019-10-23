package heapsort

func HeapSort(a []int) {
	if len(a) < 2 { return  }

	maxIndex := len(a) -1
	for i := (maxIndex) / 2; 0 <= i; i-- {
		siftDown(a, i, maxIndex)
	}
	for {
		a[0], a[maxIndex] = a[maxIndex], a[0]
		maxIndex--
		if maxIndex <= 0 { break }
		siftDown(a, 0, maxIndex)
	}
}

func siftDown(a []int, i, maxIndex int) {
	tmp := a[i]
	for j := 2*1+1; j < maxIndex; j = 2*i+1 {
		if j < maxIndex && a[j] < a[j+1] {
			j++
		}
		if a[j] <= tmp {
			 break
		}
		a[i], i = a[j], j
	}
	a[i] = tmp
}