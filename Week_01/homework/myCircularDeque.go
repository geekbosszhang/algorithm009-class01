package queue

type MyCircularDeque struct {
    head   *Node
    tail   *Node
    length  int
    count   int
}

type Node struct{
    next *Node
    prev *Node
    val  int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
    head := &Node{nil, nil, -1}
    tail := &Node{nil, nil, -1}
    head.next = tail
    tail.prev = head
    return MyCircularDeque {
        head: head,
        tail: tail,
        length: k,
        count: 0,
    }
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
    if this.IsFull() {
        return false
    }
    temp := this.head.next
    newNode := &Node {
        next: temp,
        prev: this.head,
        val: value,
    }
    temp.prev = newNode
    this.head.next = newNode
    this.count++
    return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
    if this.IsFull() {
        return false
    }
    temp := this.tail.prev
    newNode := &Node {
        next: this.tail,
        prev: temp,
        val: value,
    }
    this.tail.prev = newNode
    temp.next = newNode
    this.count++
    return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
    if this.IsEmpty() {
        return false
    }
    deleteNode := this.head.next
    this.head.next = deleteNode.next
    deleteNode.next.prev = this.head
    deleteNode.next = nil
    deleteNode.prev = nil
    this.count--
    return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
    if this.IsEmpty() {
        return false
    }
    deleteNode := this.tail.prev
    deleteNode.prev.next = this.tail
    this.tail.prev = deleteNode.prev
    deleteNode.next = nil
    deleteNode.prev = nil
    this.count--
    return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
    return this.head.next.val
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
    return this.tail.prev.val
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
    return this.head.next == this.tail
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
    return this.length == this.count
}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */