package stack

type MinStack struct {
    stack []int
    minStack []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
        stack: []int{},
        minStack: []int{},
    }
}


func (this *MinStack) Push(x int)  {
    this.stack = append(this.stack, x)
    if len(this.stack) < 1 {
        this.stack = append(this.stack, x)
    }
    if len(this.minStack) < 1 {
        this.minStack = append(this.minStack, x)
    } else {
        top := this.minStack[len(this.minStack)-1]
        if x <= top {
            this.minStack = append(this.minStack, x)
        }
    }
}


func (this *MinStack) Pop()  {
    var top int
    if len(this.stack) > 0 {
        top = this.stack[len(this.stack)-1]
        this.stack = this.stack[:len(this.stack)-1]
    }
    if len(this.minStack) > 0 {
        // pop out minStack only when the value is same with the top in stack
        if this.minStack[len(this.minStack)-1] == top {
             this.minStack = this.minStack[:len(this.minStack)-1]
        }
    }
}


func (this *MinStack) Top() int {
    if len(this.stack) > 0 {
        return this.stack[len(this.stack)-1]
    }
    return 0
}


func (this *MinStack) GetMin() int {
    if len(this.minStack) > 0 {
        return this.minStack[len(this.minStack)-1]
    }
    return 0
}