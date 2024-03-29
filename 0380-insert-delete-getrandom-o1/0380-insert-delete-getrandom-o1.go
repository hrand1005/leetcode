import (
    "math/rand"
)

type RandomizedSet struct {
    m map[int]int
    s []int
}


func Constructor() RandomizedSet {
    return RandomizedSet{
        m: make(map[int]int),
        s: []int{},
    }
}


func (this *RandomizedSet) Insert(val int) bool {
    if _, ok := this.m[val]; ok {
        return false
    }
    this.m[val] = len(this.s)
    this.s = append(this.s, val)
    return true
}


func (this *RandomizedSet) Remove(val int) bool {
    if _, ok := this.m[val]; !ok {
        return false
    }
    idx := this.m[val]
    this.s[idx] = this.s[len(this.s)-1]
    this.m[this.s[idx]] = idx
    
    delete(this.m, val)
    this.s = this.s[:len(this.s)-1]
    return true
}


func (this *RandomizedSet) GetRandom() int {
    return this.s[rand.Intn(len(this.s))]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */