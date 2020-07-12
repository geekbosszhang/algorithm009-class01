package homework

func reversePairs(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, l, r int) int {
    if l >= r {
        return 0
    }
    mid := l + (r-l) / 2
    count := mergeSort(nums, l, mid) + mergeSort(nums, mid+1, r)
    temp := make([]int, r-l+1)
    i := l
    t := l
    c := 0
    for j := mid + 1; j <= r; j++ {
        for i <= mid && nums[i] <= 2 * nums[j] {
            i++
        }
        for t <= mid && nums[t] < nums[j] {
            temp[c] = nums[t]
            t++
            c++
        }
        temp[c]=nums[j]
        count += mid - i + 1
        c++
    }

    for t<= mid {
        temp[c]=nums[t]
        c++
        t++
    }
    copy(nums[l:r+1], temp)
    return count
}