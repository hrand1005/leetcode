class Solution:
    def combinationSum2(self, candidates: List[int], target: int) -> List[List[int]]:
        self.combos = []
        candidates.sort()
        self.combo_sum_rec(candidates, target, [])
        return self.combos

    def combo_sum_rec(self, candidates: List[int], target: int, path: List[int]):
        if target == 0:
            self.combos.append(path)
            return
        
        if target < 0:
            return 
        
        for i in range(len(candidates)):
            if i > 0 and candidates[i] == candidates[i-1]:
                continue
            self.combo_sum_rec(candidates[i+1:], target-candidates[i], path + [candidates[i]])    
        return