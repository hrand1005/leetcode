type MinStack struct {
    n *node 
}

type node struct {
    prev *node
    localMin int
    val int
}


func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(val int)  {
    if this.n == nil {
        this.n = &node{
            localMin: val,
            val: val,
        }
    } else {
        this.n = &node{
            prev: this.n,
            localMin: min(this.n.localMin, val),
            val: val,
        }
    }
}


func (this *MinStack) Pop() {
    this.n = this.n.prev
}


func (this *MinStack) Top() int {
    return this.n.val
}


func (this *MinStack) GetMin() int {
    return this.n.localMin
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */