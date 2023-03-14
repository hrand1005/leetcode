/**
 * @param {string[]} strs
 * @return {string}
 */
var longestCommonPrefix = function(strs) {
    let longest = "";
    
    for (let i = 0; i < strs[0].length; i++) {
        if (strs[0].length - 1 < i) {
            return longest;
        }
        let letter = strs[0][i];
        
        for (let j = 1; j < strs.length; j++) {
            if (i > strs[j].length - 1) {
                return longest;
            }
            if (strs[j][i] !== letter) {
                return longest;
            }
        }
        
        longest += letter;
    }
    
    return longest;
};