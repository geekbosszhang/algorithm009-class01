package homework

func relativeSortArray(arr1 []int, arr2 []int) []int {
    count := make([]int, 1001)
    for _, v := range arr1 {
        count[v]++
    }
    i := 0
    for _, v2 := range arr2 {
        for count[v2] > 0 {
            arr1[i] = v2
            i++
            count[v2]--
        }
    }
    for nums1:= 0; nums1 < len(count); nums1++ {
        for count[nums1] > 0 {
            arr1[i] = nums1
            i++
            count[nums1]--
        }
    }
    return arr1
}