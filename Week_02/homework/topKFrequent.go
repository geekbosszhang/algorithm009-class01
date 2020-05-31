package homework

import "container/heap"

type FrequencyNode struct {
    Val   int
    times int
}
type heapNode []*FrequencyNode

func (h heapNode) Len() int {return len(h)}
func (h heapNode) Less(i, j int) bool {return h[i].times >  h[j].times}
func (h heapNode) Swap(i, j int) {h[i], h[j] = h[j], h[i]}

func (h *heapNode) Push(x interface{}) {
    *h = append(*h, x.(*FrequencyNode))
}

func (h *heapNode) Pop() interface{} {
    old := *h 
    n := len(old)
    x := old[n-1]
    *h = old[0:n-1]
    return x
}

func topKFrequent(nums []int, k int) (res []int) {
    hashMap := make(map[int]int, 0)
    for _, v := range nums {
        hashMap[v]++
    }
    h := &heapNode{}
    for k, v := range hashMap {
        node := &FrequencyNode{
            Val: k,
            times: v,
        }
        heap.Push(h, node)
	}
	topKLength := k 
	if len(hashMap) < k {
		topKLength = len(hashMap)
	}
    for i:= 0 ; i < topKLength; i++ {
        res = append(res, heap.Pop(h).(*FrequencyNode).Val)
    }
    return res
}