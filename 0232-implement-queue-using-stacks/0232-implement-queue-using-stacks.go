/*
type MyQueue struct {
    s1 []int
    s2 []int
    peek int
}


func Constructor() MyQueue {
    return MyQueue{
        s1: []int{},
        s2: []int{},
    }
}


func (this *MyQueue) Push(x int)  {
    if len(this.s1) == 0 {
        this.peek = x
    }
    this.s1 = append(this.s1, x)
}


func (this *MyQueue) Pop() int {
    for len(this.s1) > 1 {
        popped := this.s1[len(this.s1)-1]
        this.s1 = this.s1[:len(this.s1)-1]
        this.s2 = append(this.s2, popped)
        this.peek = popped
    }
    result := this.s1[0]
    this.s1 = this.s1[:len(this.s1)-1]
    
    for len(this.s2) > 0 {
        popped := this.s2[len(this.s2)-1]
        this.s2 = this.s2[:len(this.s2)-1]
        this.s1 = append(this.s1, popped)
    }
    return result
}


func (this *MyQueue) Peek() int {
    return this.peek
}


func (this *MyQueue) Empty() bool {
    return len(this.s1) == 0
}
*/

type MyQueue struct {
    in []int
    out []int
}


func Constructor() MyQueue {
    return MyQueue{
        in: []int{},
        out: []int{},
    }
}


func (this *MyQueue) Push(x int)  {
    this.in = append(this.in, x)
}


func (this *MyQueue) Pop() int {
    this.flush()
    popped := this.out[len(this.out)-1]
    this.out = this.out[:len(this.out)-1]
    return popped
}


func (this *MyQueue) Peek() int {
    this.flush()
    return this.out[len(this.out)-1]
}


func (this *MyQueue) Empty() bool {
    return len(this.in) == 0 && len(this.out) == 0
}

// flush dumps elements from in-stack to out-stack
// if the out-stack is empty
func (this *MyQueue) flush() {
    if len(this.out) == 0 {
        for len(this.in) > 0 {
            flushed := this.in[len(this.in)-1]
            this.in = this.in[:len(this.in)-1]
            this.out = append(this.out, flushed)
        }
    }
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */