package homework

type LRUCache struct {
    node []lruNode

    head *lruNode
    tail *lruNode

    capacity int
    used int
}

type lruNode struct {
    prev *lruNode
    next *lruNode
    key int 
    value int

    hnext *lruNode
}

const LENGTH = 100

func Constructor(capacity int) LRUCache {
    return LRUCache{
        node: make([]lruNode, LENGTH),
        head: nil,
        tail: nil,
        capacity: capacity,
        used: 0,
    }
}


func (this *LRUCache) Get(key int) int {
    if this.tail == nil {
        return -1
    }

    if tmp := this.searchNode(key); tmp!= nil {
        this.moveToTail(tmp)
        return tmp.value
    }

    return -1
}

func (this *LRUCache) searchNode(key int) *lruNode {
    if this.tail == nil {
        return nil
    }

    tmp := this.node[hash(key)].hnext
    for tmp != nil {
        if tmp.key == key {
            return tmp
        }
        tmp = tmp.hnext
    }
    return nil
} 

func (this *LRUCache) moveToTail(node *lruNode) {
    if this.tail == node {
        return
    }

    if this.head == node {
        this.head = node.next
        this.head.prev = nil
    } else {
        node.next.prev = node.prev
        node.prev.next = node.next
    }

    node.next = nil
    this.tail.next = node
    node.prev = this.tail

    this.tail = node
}

func hash(key int) int {
    return (key ^ (key >> 32)) & (LENGTH - 1)
}

func (this *LRUCache) Put(key int, value int)  {
    //1. 插入数据，数据不再LRU中
    //2. 插入数据，数据在LRU中
    //3. 插入数据不再LRU中，且已满
    if tmp := this.searchNode(key); tmp != nil {
        tmp.value = value
        this.moveToTail(tmp)
        return
    }
    this.addNode(key, value)

    if this.used > this.capacity {
        this.delNode()
    }
}

func (this *LRUCache) addNode(key int, value int) {
    newNode := &lruNode{
        key: key,
        value: value,
    }

    // why hnext here?
    tmp := &this.node[hash(key)]
    newNode.hnext = tmp.hnext
    tmp.hnext = newNode


    this.used++

    if this.tail == nil {
        this.tail, this.head = newNode, newNode
    }
    this.tail.next = newNode
    newNode.prev = this.tail
    this.tail = newNode
}

func (this *LRUCache) delNode() {
    if this.head == nil {
        return
    }

    prev := &this.node[hash(this.head.key)]
    tmp := prev.next

    for tmp != nil && tmp.key != this.head.key {
        prev = tmp
        tmp = tmp.hnext
    }
    if tmp == nil {
        return 
    }
    prev.hnext = tmp.hnext
    this.head = this.head.next
    this.head.prev = nil
    this.used--
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */