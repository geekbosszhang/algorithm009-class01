package array

import "sort"

func threeSum(nums []int) (res [][]int) {
    if len(nums) < 3 {
        return nil
    }
    // sort nums
    sort.Ints(nums)
    for k:= 0; k < len(nums) - 2; k++ {
        l := k+1
        r := len(nums) - 1
        if nums[k] > 0 || nums[k] + nums[l] > 0 {
            break
        }
        // skip duplicate k
        if k > 0 && nums[k] == nums[k-1] {
            continue
        }
        for l < r {
            if nums[k] + nums[l] + nums[r] > 0 {
                r--
            } else if nums[k] + nums[l] + nums[r] < 0 {
                l++
            } else {
                res = append(res, []int{nums[k], nums[l], nums[r]})
                l++
                r--
                // skip duplicate value
                for l < r && nums[l] == nums[l-1] {
                    l++
                }
                // skip duplicate value
                for l < r && nums[r] == nums[r+1] {
                    r--
                }
            }
        }
    }
    return res
}
