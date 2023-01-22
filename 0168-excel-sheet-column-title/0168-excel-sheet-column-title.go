func convertToTitle(columnNumber int) string {
    title := ""
    for 0 < columnNumber {
        rem := columnNumber % 26
        if rem == 0 {
            rem = 26
        }
        columnNumber = (columnNumber - rem) / 26
        title = toLetter(rem) + title
    }
    return title
}

func toLetter(x int) string {
    return string(rune(x+64))
}