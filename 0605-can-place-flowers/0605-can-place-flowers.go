func canPlaceFlowers(flowerbed []int, n int) bool {
    capacity := 0
    pos := 0
    for pos < len(flowerbed) {
        validBefore := (pos == 0) || (flowerbed[pos-1] == 0)
        validAfter := (pos == len(flowerbed) - 1) || (flowerbed[pos+1] == 0)
        if validBefore && validAfter && flowerbed[pos] == 0 {
            capacity++
            pos += 2
        } else {
            pos++
        }
    }
    return n <= capacity
}
