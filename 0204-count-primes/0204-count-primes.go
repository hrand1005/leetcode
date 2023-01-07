const (
    prime = 0
    composite = 1
)

func countPrimes(n int) int {
    if n < 3 {
        return 0
    }
    
    primeTable := make([]int, n)
    primeTable[0], primeTable[1] = composite, composite
    for i := 2; i < n; i++ {
        if primeTable[i] == prime {
            multiple := i * i
            for multiple < n {
                primeTable[multiple] = composite
                multiple += i
            }
        }
    }
    
    primeCount := 0
    for _, v := range primeTable {
        if v == prime {
            primeCount ++
        }
    }
    return primeCount
}