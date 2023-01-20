class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        freq = {}
        for n in nums:
            freq[n] = freq.get(n, 0) + 1
        
        frequencies = []
        freq_to_val = {}
        for num, freq in freq.items():
            frequencies.append(freq)
            freq_to_val[freq] = freq_to_val.get(freq, []) + [num]
        
        frequencies.sort()
        top_k_freqs = set(frequencies[-k:])
        
        result = []
        for n in top_k_freqs:
            result += freq_to_val[n]
        
        return result
        
            
            