/**
 * @param {number} x
 * @return {boolean}
 */
let isPalindrome = function(x) {
    const numString = x.toString();
    for (let i = 0; i < numString.length/2; i++) {
        if (numString[i] != numString[numString.length-1-i]) {
            return false;
        }
    }
    return true;
};