import (
    "strings"
)

func reverseVowels(s string) string {
    const vowels = "aeiouAEIOU"
    sSlice := strings.Split(s, "")
    
    left, right := 0, len(sSlice)-1
    for left < right {
        for left < right && !strings.Contains(vowels, sSlice[left]) {
            left++
        }
        for right > left && !strings.Contains(vowels, sSlice[right]) {
            right--
        }
        sSlice[left], sSlice[right] = sSlice[right], sSlice[left]
        left++
        right--
    } 
    return strings.Join(sSlice, "")
}