/*
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
*/

func canPlaceFlowers(flowerbed []int, n int) bool {
    return maxPlantable(flowerbed) >= n
}

func maxPlantable(flowerbed[]int) int {
    if len(flowerbed) == 0 {
        return 0
    }
    if len(flowerbed) == 1 {
        if flowerbed[0] == 0 {
            return 1
        }
        return 0
    }
    
    if flowerbed[0] == 0 && flowerbed[1] == 0 {
        return maxPlantable(flowerbed[2:]) + 1
    }
    if flowerbed[0] == 0 {
        return maxPlantable(flowerbed[1:])
    }
    return maxPlantable(flowerbed[2:])
}