class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        combos = []
        for i in range(len(candidates)):
            if candidates[i] == target:
                combos.append([candidates[i]])
            if candidates[i] < target:
                subcombos = self.combinationSum(candidates[i:], target-candidates[i])
                combos += [[candidates[i]] + s for s in subcombos]
        return combos