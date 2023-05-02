/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isAnagram = function(s, t) {
    let count = {}
    for (let c of s) {
        count[c] = (count[c] || 0) + 1
    }
    
    for (let c of t) {
        if (count[c] === undefined || count[c] == 0) {
            return false
        }
        count[c] = count[c] - 1
    }
    return s.length == t.length
};