type node struct {
    key int
    val int
    next *node
    prev *node
}

type LRUCache struct {
    capacity int
    m map[int]*node
    head *node
    tail *node
}


func Constructor(capacity int) LRUCache {
    head, tail := &node{}, &node{}
    head.next = tail
    tail.prev = head
    return LRUCache{
        capacity: capacity,
        m: make(map[int]*node),
        head: head,
        tail: tail,
    }
}


func (this *LRUCache) Get(key int) int {
    if n, ok := this.m[key]; ok {
        this.removeNode(n)
        this.addToTail(n)
        return n.val
    } 
    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    var newNode *node
    if n, ok := this.m[key]; ok {
        this.removeNode(n)
        n.val = value
        newNode = n
    } else {
        if len(this.m) == this.capacity {
            rem := this.m[this.head.next.key]
            this.removeNode(rem)
            delete(this.m, rem.key)
        }
        newNode = &node{
            key: key,
            val: value,
        }
    }
    this.addToTail(newNode)
    this.m[key] = newNode
}

func (this *LRUCache) removeNode(n *node) {
    prev := n.prev
    next := n.next
    prev.next = next
    next.prev = prev
}

func (this *LRUCache) addToTail(n *node) {
    prev := this.tail.prev
    prev.next = n
    n.prev = prev
    n.next = this.tail
    this.tail.prev = n
}