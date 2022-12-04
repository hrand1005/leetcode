/*
func hammingWeight(num uint32) int {
    if num == 0 {
        return 0
    } 
    
    count := 0
    for num > 1 {
        if num % 2 == 1 {
            count++
        }
        num >>= 1
    }
    
    return count + 1
}
*/

func hammingWeight(num uint32) int {
    if num == 0 {
        return 0
    } 
    
    count := 0
    for num > 1 {
        if num % 2 == 1 {
            count++
        }
        num /= 2
    }
    
    return count + 1
}