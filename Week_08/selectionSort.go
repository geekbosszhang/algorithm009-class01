package Week_08

func selectionSort(a []int) {
	for i := 0; i < len(a) - 1; i++ {
		minIndex := i
		for j := i+1; j < len(a); j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}
		a[i], a[minIndex] = a[minIndex], a[i]
	}
}