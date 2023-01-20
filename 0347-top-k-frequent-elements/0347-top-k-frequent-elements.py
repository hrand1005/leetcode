class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        freq = {}
        for n in nums:
            freq[n] = freq.get(n, 0) + 1
        
        # frequencies = []
        freq_to_val = {}
        for num, freq in freq.items():
            # frequencies.append(freq)
            freq_to_val[freq] = freq_to_val.get(freq, []) + [num]
        
        frequencies = []
        idx = len(nums)
        while len(frequencies) < k:
            if freq_to_val.get(idx) != None:
                frequencies += freq_to_val[idx]
            idx -= 1
            
        return frequencies
        
            
            