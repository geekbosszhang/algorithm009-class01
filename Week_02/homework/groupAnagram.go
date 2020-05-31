package homework

import "sort"

type bytes []byte

func (b bytes) Len() int {
    return len(b)
}

func (b bytes) Less(i, j int) bool {
    return b[i] < b[j]
}

func (b bytes) Swap(i,j int) {
    b[i], b[j] = b[j], b[i]
}

func groupAnagrams(strs []string) (res [][]string) {
    m := make(map[string]int)
    for _, v := range strs {
        sort.Sort(bytes(v))
        kBytes := string(bytes(v))
        if idx, ok := m[kBytes]; !ok {
            m[kBytes] = len(res)
            res = append(res, []string{v})
        } else {
            res[idx] = append(res[idx], v)
        }

    }
    return res
}