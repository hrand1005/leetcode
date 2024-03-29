/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

/*
type NestedIterator struct {
    index int
    flat []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    return &NestedIterator{
        index: 0,
        flat: initRecursive(nestedList),
    }
}

func initRecursive(nestedList []*NestedInteger) []int {
    flat := make([]int, 0, len(nestedList))
    for _, n := range nestedList {
        if n.IsInteger() {
            flat = append(flat, n.GetInteger())
        } else {
            flat = append(flat, initRecursive(n.GetList())...)
        }
    }
    return flat
}

func (this *NestedIterator) Next() int {
    next := this.flat[this.index]
    this.index++
    return next
}

func (this *NestedIterator) HasNext() bool {
    if this.index < len(this.flat) {
        return true
    }
    return false
}
*/

type NestedIterator struct {
    head int
    stack []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    return &NestedIterator{
        stack: nestedList,
    }
}

func (this *NestedIterator) Next() int {
    return this.head
}

func (this *NestedIterator) HasNext() bool {
    for len(this.stack) != 0 {
        nest := this.stack[0]
        this.stack = this.stack[1:]
        if nest.IsInteger() {
            this.head = nest.GetInteger()
            return true
        }
        this.stack = append(nest.GetList(), this.stack...)
    }
    return false
}