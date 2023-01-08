class Trie:

    def __init__(self):
        self.words = {}
        self.prefixes = {}

    def insert(self, word: str) -> None:
        self.words[word] = True
        for i in range(1, len(word)):
            self.prefixes[word[:i]] = True
        
    def search(self, word: str) -> bool:
        return self.words.get(word, False)
        
    def startsWith(self, prefix: str) -> bool:
        return self.prefixes.get(prefix, False) or self.words.get(prefix, False)
        
