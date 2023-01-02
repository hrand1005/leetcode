/*
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
*/

type node struct {
    prev *node
    val int
    minVal int
}

type MinStack struct {
    current *node
}


func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(val int)  {
    if this.current == nil {
        this.current = &node{
            val: val,
            minVal: val,
        }
    } else {
        minVal := min(this.current.minVal, val)
        this.current = &node{
            val: val,
            minVal: minVal,
            prev: this.current,
        }
    }
}


func (this *MinStack) Pop()  {
    this.current = this.current.prev
}


func (this *MinStack) Top() int {
    return this.current.val
}


func (this *MinStack) GetMin() int {
    return this.current.minVal
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}