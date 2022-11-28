func romanToInt(s string) int {
    intVal := map[rune]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }
    
    prev := ' '
    sum := 0
    for _, v := range s {
        sum += intVal[v]
        if intVal[prev] < intVal[v] {
            sum -= 2*intVal[prev]
        }
        prev = v
    }
    
    return sum
}