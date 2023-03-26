/**
 * @param {number[]} digits
 * @return {number[]}
 */
var plusOne = function(digits) {
    if (digits.length === 0) {
        return [1]
    }
    
    if (digits[digits.length-1] + 1 > 9) {
        let ret = plusOne(digits.slice(0, -1))
        ret.push(digits[digits.length-1]-9)
        return ret
    }
    
    let ret = digits.slice(0, -1)
    ret.push(digits[digits.length-1]+1)
    return ret
};