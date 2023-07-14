func decodeString(s string) string {
    decoded := ""
    for _, c := range s {
        if c == ']' {
            decoded = decodeLast(decoded)
        } else {
            decoded += string(c)
        }
    }
    return decoded
}

func decodeLast(encoding string) string {
    start := strings.LastIndex(encoding, "[")
    sub := encoding[start+1:]
    
    end := start
    start--
    
    for start >= 0 && unicode.IsDigit(rune(encoding[start])) {
        start--
    }
    rep, _ := strconv.Atoi(string(encoding[start+1:end]))
    return encoding[:start+1] + strings.Repeat(sub, rep)
}