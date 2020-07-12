学习笔记

超哥直播笔记
1. 爬楼梯问题
时间复杂度O(n), O(1)空间复杂度
```
func climbStairs(n int) int {
    if n < 3 {
        return n
    }
    n1 := 1
    n2 := 2

    for i:= 3; i <= n; i++ {
        n2, n1 = n1+n2, n2
    }
    return n2
}
```
扩展打印出爬楼梯的路径 
```
def climbStairs(self, n):
        steps = [1, 2]
        self._dfs(n, [], steps)
    
    def _dfs(self, n, res, steps):
        if n == 0:
            print(res)
            return

        for step in steps:
            if n >= step:
                self._dfs(n-step, res + [step], steps)
```
2. 有效的括号
借助栈的实现
```
func isValid(s string) bool {
    if (len(s)%2)&1 == 1 {
        return false
    }
    parentMap := map[rune]rune{
        '{': '}',
        '[': ']',
        '(': ')',
    }
    wordStack := make([]rune, 0)
    for _, v := range s {
        if value, ok := parentMap[v]; ok {
            wordStack = append(wordStack, value)
        } else {
            if len(wordStack) == 0 {
                return false
            }
            if v == wordStack[len(wordStack)-1] {
                wordStack = wordStack[:len(wordStack)-1]
            } else {
                return false
            }
        }
    }
    return len(wordStack) == 0
}
```
3. Pow(x, n)
思路 借助递归和二分
```
func myPow(x float64, n int) float64 {
    if n < 0 {
        return 1 / quickPow(x, -n)
    }
    return quickPow(x, n)
}

func quickPow(x float64, n int) float64{
    // terminater
    if n == 0 {
        return 1
    }
    // process current logic
    half := quickPow(x, n>>1)
    if n&1 == 1 {
       return x * half * half
    }
    return half * half
}
```
4. 路径总和
从根节点到叶子节点
```
func hasPathSum(root *TreeNode, sum int) bool {
    if root == nil {
        return false
    }
    if root.Left == nil && root.Right == nil {
        return sum == root.Val
    }
    return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum - root.Val)
}
```
变形题，不一定从根节点开始
```
func hasPathSum(root *TreeNode, sum int) bool {
    if root == nil {
        return false
    }
    if root.Left == nil && root.Right == nil {
        return sum == root.Val
    }
    return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum - root.Val) || hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}
```

排序笔记

基本排序
1. 选择排序
每次找最小值(O(n))，然后放到待排序数组的起始位置
2. 插入排序
从前到后构建有序序列；对于未排序数据，在已排序序列中从后向前扫描，插入到相应的位置 
即便可以二分查找对应的位置，但是数组挪动元素仍需要O(n),外层遍历也需要O(n)
3. 冒泡排序 
嵌套循环，每次查看相邻的元素如果逆序，则交换，（不常用）
循环一层能选择出最大的元素，与选择排序相反

高级排序
1. 快速排序
基于分治的思想
数组取标杆pivot,将小元素放pivot左边，大元素放右侧，然后依次对左边和右边的子数组继续快排，最后达到整个序列有序
关键是pivot元素的选取 （不开辟新的内存空间怎么partition）？？？

2. 归并排序
把长度为n的输入序列分成两个长度为n/2的子序列
对这两个子序列分别采用归并排序
将两个排序好的合并成一个

```
func mergeSort(array []int, left, right int) {
    if right <= left {
        return
    }
    mid := (left+right) >> 1

    mergeSort(array, left, mid)
    mergeSort(array, mid+1, right)
    merge(array, left, mid, right)
}
```

merge合并两个有序数组 或者有序链表

快排VS归并

3. 堆排序 - 堆插入O(logN), 取最大最小 O(1)
数组元素依次建立小顶堆
pop出堆顶元素

实现代码在Week_08下面
