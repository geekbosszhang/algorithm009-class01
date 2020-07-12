package Week_08

import "fmt"

func BubbleSort(a []int) {
	if len(a) <= 1 {
		return
	}
	for i := 0; i < len(a) - 1;i++ {
		for j := 0; j < len(a) - i - 1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		} 
	}
}

func main() {
	a := []int{2,3,4,1,5,7,6,19,22,10}
	BubbleSort(a)
	fmt.Println("a is", a)
}