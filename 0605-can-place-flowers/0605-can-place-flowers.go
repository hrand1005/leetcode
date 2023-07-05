func canPlaceFlowers(flowerbed []int, n int) bool {
    capacity := 0
    pos := 0
    for pos < len(flowerbed) {
        if canPlant(flowerbed, pos) {
            capacity++
            pos += 2
        } else {
            pos++
        }
    }
    return n <= capacity
}

func canPlant(flowerbed []int, pos int) bool {
    validAfter := (pos == len(flowerbed) - 1) || (flowerbed[pos+1] == 0)
    validBefore := (pos == 0) || (flowerbed[pos-1] == 0)
    return validBefore && validAfter && flowerbed[pos] == 0
}