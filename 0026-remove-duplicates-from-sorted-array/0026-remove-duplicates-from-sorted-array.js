/**
 * @param {number[]} nums
 * @return {number}
 */
var removeDuplicates = function(nums) {
    let insertAt = 1
    let currentVal = nums[0]
    
    for (let i = 1; i < nums.length; i++) {
        if (nums[i] !== currentVal) {
            nums[insertAt] = nums[i]
            currentVal = nums[i]
            insertAt++
        }
    }
    
    return insertAt
};