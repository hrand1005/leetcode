class Trie:

    def __init__(self):
        self.words = {}
        self.prefixes = {}

    def insert(self, word: str) -> None:
        self.words[word] = True
        for i in range(1, len(word)+1):
            self.prefixes[word[:i]] = True
        
    def search(self, word: str) -> bool:
        return self.words.get(word, False)
        
    def startsWith(self, prefix: str) -> bool:
        return self.prefixes.get(prefix, False)
        


# Your Trie object will be instantiated and called as such:
# obj = Trie()
# obj.insert(word)
# param_2 = obj.search(word)
# param_3 = obj.startsWith(prefix)