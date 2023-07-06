func compress(chars []byte) int {
    char := byte(' ')
    count := 0
    writeIdx := 0
    
    for _, c := range chars {
        if c == char {
            count++
        } else {
            if count == 1 {
                chars[writeIdx] = char
                writeIdx++
            } else if count > 1 {
                chars[writeIdx] = char
                writeIdx++
                countString := strconv.Itoa(count)
                for i := 0; i < len(countString); i++ {
                    chars[writeIdx] = countString[i]
                    writeIdx++
                }
            }
            char = c
            count = 1
        }
    }
    if count == 1 {
        chars[writeIdx] = char
        writeIdx++
    } else if count > 1 {
        chars[writeIdx] = char
        writeIdx++
        countString := strconv.Itoa(count)
        for i := 0; i < len(countString); i++ {
            chars[writeIdx] = countString[i]
            writeIdx++
        }
    }
    return writeIdx
}