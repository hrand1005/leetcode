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


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */