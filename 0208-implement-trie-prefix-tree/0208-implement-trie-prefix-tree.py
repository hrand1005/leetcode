"""
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
"""

class Trie:

    def __init__(self):
        self.trie = {}

    def insert(self, word: str) -> None:
        trie = self.trie
        for letter in word:
            if letter not in trie:
                trie[letter] = {}
            trie = trie[letter]
        trie["\0"] = True
        
    def search(self, word: str) -> bool:
        trie = self.trie
        for letter in word:
            if letter not in trie:
                return False
            trie = trie[letter]
        return "\0" in trie
        
    def startsWith(self, prefix: str) -> bool:
        trie = self.trie
        for letter in prefix:
            if letter not in trie:
                return False
            trie = trie[letter]
        return True