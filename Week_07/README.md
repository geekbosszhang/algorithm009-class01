Trie 树的学习总结

Trie树又称字典树，用来解决在一组字符串集合中快速查找某个字符串的问题。
用于统计和排序大量的字符串，用作文本词频的统计，能够最大限度的减少无谓的字符串比较，查询效率比Hash高。
*** 特点是 ***
1. 节点本身不存完整单词
2. 从根节点到某一节点路径上经过的字符连接起来为该单词
3. 所有子节点路径都不相同

但是缺点是空间消耗大，相当于用空间换时间，查找的次数与单词长度有关
Trie树经典的使用场景是搜索引擎的关键词提示功能

***  Trie树与散列表和红黑树的比较 ***
最大的优势就是做前缀匹配的字符串搜索
而动态集合数据的从查找可以用更合适的散列表和红黑树来替代

相关习题一
1. Trie树的实现

```
class TrieNode:
# Initialize your data structure here.
def __init__(self):
    self.children = collections.defaultdict(TrieNode)
    self.is_word = False

class Trie:

def __init__(self):
    self.root = TrieNode()

def insert(self, word):
    current = self.root
    for letter in word:
        current = current.children[letter]
    current.is_word = True

def search(self, word):
    current = self.root
    for letter in word:
        current = current.children.get(letter)
        if current is None:
            return False
    return current.is_word

def startsWith(self, prefix):
    current = self.root
    for letter in prefix:
        current = current.children.get(letter)
        if current is None:
            return False
    return True
```

2. 单词搜索II

Refered https://leetcode.com/problems/word-search-ii/discuss/59780/Java-15ms-Easiest-Solution-(100.00)

时间复杂度分析
```
N = num of rows
M = num of columns
X = number of words in dictionary
Y = length of longest word in dictionary

Time complexity of using the trie tree

We are doing a 4-child DFS traversal.
The worst case depth in this case is the minimum of [traversing the entire board, traversing until we hit the end of a word].
-> O of each traversal is4^(min(Y, NM))
Note: min accounts for the case where we have words in the dictionary longer than the number of characters in the board itself.
We are doing this N*M times since we need to check for words starting at each position in the board
-> O(4^(min(Y, NM))*NM)
```
Golang的代码实现
```
type Trie struct {
    data string
    children [26]*Trie
    isEndingChar bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    for _, v := range word {
        index := v - 'a'
        if this.children[index] == nil {
            newNode := new(Trie)
            this.children[index] = newNode
        }
        this = this.children[index]
    }
    this.data = word
    this.isEndingChar = true
}

func findWords(board [][]byte, words []string) (res []string) {
    // 遍历word， 建立Trie树
    trie := new(Trie)
    for _, v := range words {
        trie.Insert(v)
    }
    // Board DFS search
    for i:= 0; i < len(board); i++  {
        for j := 0; j < len(board[0]); j++ {
            existRecursive(board, i, j, trie, &res);
        }
    }
    return res
}

func existRecursive(board [][]byte, row, col int, trie *Trie, res *[]string) {
    if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) {
        return
    }
    cur := board[row][col]
    if cur == '$' || trie.children[cur - 'a'] == nil {
        return
    }
    trie = trie.children[cur - 'a']
    if trie.data != "" {
        *res = append(*res, trie.data)
        trie.data = ""
    }
    temp := board[row][col]
    board[row][col] = '$'
    existRecursive(board, row-1, col, trie, res)
    existRecursive(board, row + 1, col, trie, res)
    existRecursive(board, row, col-1, trie, res)
    existRecursive(board, row, col+1, trie, res)
    board[row][col] = temp
}
```

*** 并查集 *** 

```
def parent(self, p, i):
    root = i
    while p[root] != root:
        root = p[root]
    
    #路径压缩
    while p[i] != i:
        x=i;i=p[i];p[x]=root
    
    return root
```
关于 被围绕的区域 思路：
首先我们把每个节点各作为一类，用它的行数和列数生成一个 id 标识该类。
```
int node(int i, int j) {
    return i * cols + j;
}
```
然后遍历每个 O节点，和它上下左右的节点进行合并即可。

如果当前节点是边界的 O,就把它和 dummy 节点（一个在所有节点外的节点）合并。最后就会把所有连通到边界的 o 节点和 dummy 节点合为了一类。

最后我们只需要判断，每一个 o 节点是否和 dummy 节点是不是一类即可。


*** 剪枝 ***
参考例题为判断有效数独和solve数独

*** 双向DFS *** 
关于ladderLength使用双向BFS算法
思路
1. 把开始和结束的节点都加入到两个queue里面
2. 每次选择一个queue，然后遍历一层，如果遍历的下一层有对方的节点，就退出
3. 选择另外一个queue，然后遍历一层，如果遍历的下一层有对方的节点，就退出
4. 如果任何一个queue都空了，如果还没有遇到就结束

双向BFS每次遍历还是要完整的做一层遍历，不能之遍历一个元素。
优化的点是每次从len(queue)小的开始

JAVA 版代码
```
class Solution {
    public int ladderLength(String beginWord, String endWord, List<String> wordList) {
        if(!wordList.contains(endWord)) return 0;
        
        Queue<String> queue1 = new LinkedList<>();
        queue1.add(beginWord);
        Queue<String> queue2 = new LinkedList<>();
        queue2.add(endWord);
        
        Set<String> visited = new HashSet<>();
        visited.add(endWord);
        
        int steps = 1;
        while(queue1.size() > 0 && queue2.size() > 0) {
            // always start from smaller number of queue 
            if(queue1.size() > queue2.size()) {
                Queue<String> temp = queue1;
                queue1 = queue2;
                queue2 = temp;
            }
            
            Queue<String> nextQueue = new LinkedList<>();
            while(!queue1.isEmpty()) {
                String cur = queue1.poll();
                for(String word: wordList) {
                    if(valid(cur, word)) {
                        if(queue2.contains(word)) {
                            return steps+1;
                        }
                        
                        if(!visited.contains(word)) {
                            nextQueue.add(word);
                            visited.add(word);                            
                        }
                    }
                }
            }
            queue1 = nextQueue;
            steps++;
        }
        return 0;
    }
    
    boolean valid(String a, String b) {
        int diff = 0;
        for(int i = 0; i < a.length(); ++i) {
            if(a.charAt(i) != b.charAt(i)) {
                diff++;
                if(diff >= 2) {
                    return false;
                }
            }
        }
        return true;
    }
}
```