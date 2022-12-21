LETTER_TO_NUMBER = {
     "1":"A",
     "2":"B",
     "3":"C",
     "4":"D",
     "5":"E",
     "6":"F",
     "7":"G",
     "8":"H",
     "9":"I",
     "10":"J",
     "11":"K",
     "12":"L",
     "13":"M",
     "14":"N",
     "15":"O",
     "16":"P",
     "17":"Q",
     "18":"R",
     "19":"S",
     "20":"T",
     "21":"U",
     "22":"V",
     "23":"W",
     "24":"X",
     "25":"Y",
     "26":"Z",
}
"""
class Solution:
    def numDecodings(self, s: str) -> int:
        letter_to_number = {
             "1":"A",
             "2":"B",
             "3":"C",
             "4":"D",
             "5":"E",
             "6":"F",
             "7":"G",
             "8":"H",
             "9":"I",
             "10":"J",
             "11":"K",
             "12":"L",
             "13":"M",
             "14":"N",
             "15":"O",
             "16":"P",
             "17":"Q",
             "18":"R",
             "19":"S",
             "20":"T",
             "21":"U",
             "22":"V",
             "23":"W",
             "24":"X",
             "25":"Y",
             "26":"Z",
        }
        
        ways = 0
        stack = [s]
        while stack:
            word = stack.pop()
            if len(word) == 0:
                ways += 1
                continue
                
            if letter_to_number.get(word[0]) != None:
                stack.append(word[1:])
            if len(word) >= 2 and letter_to_number.get(word[:2]) != None:
                stack.append(word[2:])
            
        return ways
"""

class Solution:
    def numDecodings(self, s: str) -> int:
        if s[0] == "0":
            return 0
        
        if len(s) == 1:
            return 1
        
        all_ways = [0] * (len(s)+1)
        
        all_ways[0] = 1
        if s[1] != "0":
            all_ways[1] = 1
        
        for i in range(2, len(s)+1):
            if s[i-1] != "0":
                all_ways[i] += all_ways[i-1]
            if LETTER_TO_NUMBER.get(s[i-2:i]) != None:
                all_ways[i] += all_ways[i-2]
        
        return all_ways[-1]
        