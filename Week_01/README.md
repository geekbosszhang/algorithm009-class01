Week1 学习笔记
切题四件套
1. Clarification 搞清楚问题是什么
2. Possible solutions 比较不同方法的时间和空间复杂度，能够选择更优化的解法
3. Coding
4. Test case (include corner case)
刷题五遍
1. 5分钟读题+思考，如果实在没思路可以直接看解法，多看Top的解法，比较优劣，能够记住较好的解法
2. 自己重头写一遍，能够提交通过多有的case，且观察时间和空间的performance，多看不同的解法进行优化
3. 24小时后，尤其是对于不熟练和起初没思路的需要再练习一遍
4. 一周后再重复的刷一遍
5. 面试前提前1周左右再刷一遍

chunk it up + 刻意练习

记住常用的快捷键能够提升效率，包括常用的fn+del能够删除后面的word，command+left/right 行首和行尾，option+left/right 一个单词一个单词的跳, shift+command可以多选等。

常见的DS包括
1. 一维数据结构 （Array，Linklist)
栈、队列、双端队列（deque), 集合，映射（map/hash)
2. 二维数据结构
树、图、二叉搜索树、堆heap, 并查集、字典树Trie
3. 特殊数据结构
位运算、布隆过滤器、LRU cache

常见的算法包括
1. if-else,switch 条件分支
2. for, while, loop 循环分支
3. 递归
4. 搜索 如深度优先搜索（DFS）、广度优先搜索（BFS)等
5. 动态规划
6. 二分查找
7. 贪心

对于数组和链表
关于动态数组的扩容和缩荣，可以参考golang中关于数组和slice切片的源码实现
对于链表的操作要重复多写几遍

对于栈和队列，在工程上双端队列更为更为常用
看了Python的collection.py关于deque的实现
https://docs.python.org/2/library/collections.html
https://bitbucket.org/pypy/pypy/src/default/lib_pypy/_collections.py
还有Golang的简单的关于双端队列的实现
https://play.golang.org/p/i1BGlRsP19
https://godoc.org/gopkg.in/karalabe/cookiejar.v1/collections/deque

优先队列（priority queue）在Python里面是借助heap实现的, 插入时间复杂度为O(1), 取出的时间复杂度为O(nlogn)，因为优先队列是按照元素的优先级顺序取出，同时优先队列适用于求解第K大（小）元素。
https://docs.python.org/2/library/heapq.html

双指针问题
快慢指针主要解决链表中的问题，比如典型的判定链表中是否包含环；
左右指针主要解决数组（或者字符串）中的问题，比如二分查找。
