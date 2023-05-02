/**
 * @param {number[]} nums
 * @return {boolean}
 */
var containsDuplicate = function(nums) {
    occ = new Map()
    for (num of nums) {
        if (occ.has(num)) {
            return true
        }
        occ.set(num, true)
    }
    return false
};