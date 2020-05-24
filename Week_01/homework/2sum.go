package array

func twoSum(nums []int, target int) (res []int) {
    hashMap := map[int]int{}
    for i, v := range nums {
        expected := target - v
        if _, ok := hashMap[expected]; ok {
            res = append(res, hashMap[expected])
            res = append(res, i)
        }
        hashMap[v] = i
    }
    return res
}