func convert(s string, numRows int) string {
    rows := make([]strings.Builder, numRows+1)
    
    idx := 0
    for idx < len(s) {
        for i := 0; i < numRows && idx < len(s); i++ {
            rows[i].WriteByte(s[idx])
            idx++
        }
        for j := numRows-2; 0 < j && idx < len(s); j-- {
            rows[j].WriteByte(s[idx])
            idx++
        }
    }
    
    var zigzag strings.Builder
    for _, r := range rows {
        zigzag.WriteString(r.String())
    }
    return zigzag.String()
}