package Week_08

func heapSort(a []int) {
	buildHeap(a)

	for i := len(a) - 1; i > 0; i-- {
		a[0], a[i] = a[i], a[0]
		heapify(a, i, 0)
	}
}

func buildHeap(a []int) {
	for i := len(a)/ 2 - 1; i >= 0; i-- {
		heapify(a, len(a), i)
	}
}

func heapify(a []int, length, i int) {
	left := 2 * i + 1
	right := 2 * i + 2
	largest := i

	if left < length && a[left] > a[largest] {
		largest = left
	}

	if right < length && a[right] > a[largest] {
		largest = right
	}

	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		heapify(a, length, largest)
	}
}