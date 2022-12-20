"""
class Node:
    def __init__(self, coord, visited, idx):
        self.coord = coord
        self.visited = visited
        self.idx = idx

class Solution:
    def exist(self, board: List[List[str]], word: str) -> bool:
        if self.is_impossible(board, word):
            return False
        
        stack = self.find_starts(board, word)
        while stack:
            node = stack.pop()
            new_visited = node.visited.copy()
            new_visited[node.coord] = True
            
            if node.idx == len(word) - 1:
                return True
            
            neighbors = self.get_neighbors(node.coord, board)
            for n in neighbors:
                if n not in new_visited:
                    if board[n[0]][n[1]] == word[node.idx+1]:
                        stack.append(
                            Node(n, new_visited, node.idx + 1)
                        )
                        
        return False
            
    def is_impossible(self, board: List[List[str]], word: str) -> bool:
        letter_count = {}
        for i in range(len(board)):
            for j in range(len(board[i])):
                if letter_count.get(board[i][j]) == None:
                    letter_count[board[i][j]] = 1
                else:    
                    letter_count[board[i][j]] += 1
        
        for c in word:
            if letter_count.get(c, 0) == 0:
                return True
            else:
                letter_count[c] -= 1
        
        return False
    
    def find_starts(self, board: List[List[str]], word: str) -> List[Node]:
        starts = []
        for i in range(len(board)):
            for j in range(len(board[i])):
                if board[i][j] == word[0]:
                    starts.append(
                        Node((i, j), {}, 0)
                    )
        return starts           

    def get_neighbors(self, coord: tuple[int, int], board: List[List[str]]) -> List[tuple[int, int]]:
        neighbors = []
        if coord[0] > 0:
            neighbors.append((coord[0]-1, coord[1]))
        if coord[0] < len(board)-1: 
            neighbors.append((coord[0]+1, coord[1]))
        if coord[1] > 0:
            neighbors.append((coord[0], coord[1]-1))
        if coord[1] < len(board[coord[0]]) - 1:
            neighbors.append((coord[0], coord[1]+1))
            
        return neighbors
"""

class Solution:
    def exist(self, board: List[List[str]], word: str) -> bool:
        if not self.possible_path(board, word):
            return False
        
        for i in range(len(board)):
            for j in range(len(board[i])):
                if self.dfs((i, j), board, word):
                    return True
                
        return False            
    
    
    def possible_path(self, board: List[List[str]], word: str) -> bool:
        letter_count = {}
        for i in range(len(board)):
            for j in range(len(board[i])):
                if letter_count.get(board[i][j]) == None:
                    letter_count[board[i][j]]= 1
                else:
                    letter_count[board[i][j]] += 1
        
        for letter in word:
            if letter_count.get(letter, 0) == 0:
                return False
            letter_count[letter] -= 1
        
        return True
    
    
    def dfs(self, coord: tuple[int, int], board: List[List[str]], word: str) -> bool:
        if len(word) == 0:
            return True
        
        i, j = coord[0], coord[1]
        if len(board)-1 < i or i < 0:
            return False
        if len(board[i])-1 < j or j < 0:
            return False
        if board[i][j] != word[0]:
            return False
        
        letter = board[i][j]
        board[i][j] = ""
        
        if self.dfs((i-1, j), board, word[1:]) or \
           self.dfs((i+1, j), board, word[1:]) or \
           self.dfs((i, j-1), board, word[1:]) or \
           self.dfs((i, j+1), board, word[1:]):
            return True
        
        board[i][j] = letter
        
        return False