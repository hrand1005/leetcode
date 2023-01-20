class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        occ = {}
        for n in nums:
            occ[n] = occ.get(n, 0) + 1
        
        occ_list = []
        occ_to_val = {}
        for value, occurrences in occ.items():
            occ_list.append(occurrences)
            occ_to_val[occurrences] = occ_to_val.get(occurrences, []) + [value]
        
        occ_list.sort()
        top_k_occurrences = set(occ_list[-k:])
        
        result = []
        for n in top_k_occurrences:
            result += occ_to_val[n]
        
        return result
        
            
            