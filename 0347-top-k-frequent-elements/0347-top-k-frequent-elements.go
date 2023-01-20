func topKFrequent(nums []int, k int) []int {
    valToFreq := make(map[int]int)
    for _, n := range nums {
        valToFreq[n]++
    }
    
    freqToVal := make(map[int][]int)
    for k, v := range valToFreq {
        freqToVal[v] = append(freqToVal[v], k)
    }
    
    freq := len(nums)
    topK := make([]int, 0, k)
    for len(topK) < k {
        topK = append(topK, freqToVal[freq]...)
        freq--
    }
    
    return topK
}