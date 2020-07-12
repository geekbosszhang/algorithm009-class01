package Week_08

func quickSort(a []int, begin, end int) {
	if begin >= end {
		return
	}
	pivot := partition(a, begin, end)
	quickSort(a, begin, pivot-1)
	quickSort(a, pivot+1, end)
}

func partition(a []int, begin, end int) int{
	pivot := end
	count := begin
	for i := begin; i < end; i++ {
		if a[i] < a[pivot] {
			a[i], a[count] = a[count], a[i]
			count++
		}
	}
	a[pivot], a[count] =  a[count], a[pivot]
	return count                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             	
}