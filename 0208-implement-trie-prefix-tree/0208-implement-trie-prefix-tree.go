type Trie struct {
    words map[string]bool
    prefixes map[string]bool
}


func Constructor() Trie {
    return Trie{
        words: make(map[string]bool),
        prefixes: make(map[string]bool),
    }
}


func (this *Trie) Insert(word string)  {
    this.words[word] = true
    for i := 1; i < len(word)+1; i++ {
        this.prefixes[word[:i]] = true
    }
}


func (this *Trie) Search(word string) bool {
    return this.words[word]
}


func (this *Trie) StartsWith(prefix string) bool {
    return this.prefixes[prefix]
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */