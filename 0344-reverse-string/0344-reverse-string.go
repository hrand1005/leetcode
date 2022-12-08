func reverseString(s []byte)  {
    for i := 0; i < len(s)/2; i++ {
        j := len(s) - 1 - i
        s[i], s[j] = s[j], s[i]
    }
}