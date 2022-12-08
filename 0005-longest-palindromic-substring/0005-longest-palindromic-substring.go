/*
func longestPalindrome(s string) string {
    maxPal := string(s[0])

    for i := 1; i < len(s); i++ {
        odd := longestPal(i, i, s)
        even := longestPal(i-1, i, s)
        maxPal = longestString(odd, even, maxPal)
    }

    return maxPal
}

// recursive version
func longestPal(l, r int, s string) string {
    if s[l] != s[r] {
        return s[l+1:r]
    }

    if l == 0 || r == len(s) - 1 {
        return s[l:r+1]
    }

    return longestPal(l-1, r+1, s)
}

// loop version
func longestPal(l, r int, s string) string {
    for l >= 0 && r < len(s) && s[l] == s[r] {
        l--
        r++
    }

    return s[l+1:r]
}

func longestString(s1, s2, s3 string) string {
    if len(s1) > len(s2) && len(s1) > len(s3) {
        return s1
    } 

    if len(s2) > len(s3) {
        return s2
    }

    return s3
}
*/

func longestPalindrome(s string) string {
    // create a lookup table such that for some
    // start index i and some end index at j,
    // palTable[i][j] = True if and only if s[i:j+1]
    // is a palindrome
    palTable := make([][]bool, len(s), len(s))
    for i := 0; i < len(s); i++ {
        palTable[i] = make([]bool, len(s), len(s))
        palTable[i][i] = true
    }

    // let the initial longest palindrome be an arbitrary
    // character in the string
    maxPal := string(s[0])

    // we want to iterate in such a way that we check the 
    // smaller substrings first, since checks on larger
    // strings will look them up in our dynamically 
    // populated table
    for i := len(s) - 1; i >= 0; i-- {
        for j := i+1; j < len(s); j++ {
            // check if s[i:j+1] could be a palindrome
            if s[i] == s[j] {
                // given that the outer characters match, 
                // there are two possible conditions that could
                // be met to prove that s[i:j+1] is a palindrome:
                //
                // 1. There is at most 1 character between i and j
                // 2. The substring s[i+1:j] is a palindrome, as 
                //    confirmed by our palTable
                if j - i < 2 || palTable[i+1][j-1] {
                    // record the palindrome in our table
                    palTable[i][j] = true
                    // update the longest palindrome if necessary
                    pal := string(s[i:j+1])
                    if len(pal) > len(maxPal) {
                        maxPal = pal
                    }
                }
            }
        }
    }

    return maxPal
}