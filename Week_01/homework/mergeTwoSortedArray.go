package array

func merge(nums1 []int, m int, nums2 []int, n int)  {
    aMaxIndex := m - 1
    bMaxIndex := n - 1
    mergeMaxIndex := m + n - 1
    for aMaxIndex >= 0 && bMaxIndex >= 0 {
        if nums1[aMaxIndex] > nums2[bMaxIndex] {
            nums1[mergeMaxIndex] = nums1[aMaxIndex]
            aMaxIndex--         
        } else {
            nums1[mergeMaxIndex] = nums2[bMaxIndex]
            bMaxIndex--
        }
        mergeMaxIndex--
    }
    for bMaxIndex >= 0 {
        nums1[mergeMaxIndex] = nums2[bMaxIndex]
        bMaxIndex--
        mergeMaxIndex--
    }
    return
}