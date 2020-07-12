package Week_08

func mergeSort(a []int, begin, end int) {
	if begin >= end {
		return
	}
	mid := (begin + end) >> 1
	mergeSort(a, begin, mid)
	mergeSort(a, mid+1, end)
	merge(a, begin, mid, end)
}

func merge(a []int, begin, mid, end int) {
	temp := make([]int, end - begin + 1)
	i := begin
	j:= mid+1
	k := 0 
	for ; i <= mid  && j <= end ; k++{
		if a[i] <= a[j] {
			temp[k] = a[i]
			i++
		} else {
			temp[k] = a[j]
			j++
		}
	}
	for ; i <= mid; i++ {
		temp[k] = a[i]
		k++
	}
	for ; j <= end; j++ {
		temp[k] = a[j]
		k++
	}
	copy(a[begin:end+1], temp)
}