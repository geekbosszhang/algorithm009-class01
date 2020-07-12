package Week_08

func insertionSort(a []int) {
	if len(a) <= 1 {
		return
	}
	for i:= 1; i < len(a); i++ {
		value := a[i]
		j := i - 1
		for ; j >= 0; j--{
			if a[j] > value {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = value
	}
}