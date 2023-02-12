/*
// USING 2 QUEUES
type MyStack struct {
    q1 []int
    q2 []int
    top int
}


func Constructor() MyStack {
    return MyStack{
        q1: make([]int, 0),
        q2: make([]int, 0),
    }
}


func (this *MyStack) Push(x int)  {
    this.q1 = append(this.q1, x)
    this.top = x
}


func (this *MyStack) Pop() int {
    n := len(this.q1)
    for i := 0; i < n-1; i++ {
        this.top = this.q1[0]
        this.q1 = this.q1[1:]
        this.q2 = append(this.q2, this.top)
    }
    pop := this.q1[0]
    this.q1 = this.q1[1:]
    this.q1, this.q2 = this.q2, this.q1
    return pop
}


func (this *MyStack) Top() int {
    return this.top
}


func (this *MyStack) Empty() bool {
    return len(this.q1) == 0
}
*/

type MyStack struct {
    queue []int
    top int
}


func Constructor() MyStack {
    return MyStack{
        queue: make([]int, 0),
    }
}


func (this *MyStack) Push(x int)  {
    this.queue = append(this.queue, x)
    this.top = x
}


func (this *MyStack) Pop() int {
    n := len(this.queue)
    for i := 0; i < n-1; i++ {
        this.top = this.queue[0]
        this.queue = this.queue[1:]
        this.queue = append(this.queue, this.top)
    }
    pop := this.queue[0]
    this.queue = this.queue[1:]
    return pop
}


func (this *MyStack) Top() int {
    return this.top
}


func (this *MyStack) Empty() bool {
    return len(this.queue) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */