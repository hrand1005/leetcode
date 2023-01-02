type MinStack struct {
    // holds pairs of {min, val}
    stack [][]int
}


func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(val int)  {
    if len(this.stack) == 0 {
        this.stack = [][]int{ {val, val} }
    } else {
        minVal := min(this.stack[len(this.stack)-1][0], val)
        this.stack = append(this.stack, []int{minVal, val})
    }
}


func (this *MinStack) Pop()  {
    this.stack = this.stack[:len(this.stack)-1]
}


func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1][1]
}


func (this *MinStack) GetMin() int {
    return this.stack[len(this.stack)-1][0]
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