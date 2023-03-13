/**
 * @param {number} x
 * @return {boolean}
 */
/*
let isPalindrome = function(x) {
    const numString = x.toString();
    for (let i = 0; i < numString.length/2; i++) {
        if (numString[i] != numString[numString.length-1-i]) {
            return false;
        }
    }
    return true;
};
*/

let isPalindrome = function(x) {
    if (x < 0) {
        return false;
    }
    let original = x;
    let reverse = 0;
    while (x != 0) {
        reverse *= 10;
        reverse += x % 10;
        x = Math.floor(x/10);
    }
    return reverse == original;
}