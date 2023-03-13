/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
let twoSum = function(nums, target) {
    const valToIdx = new Map();
    for (let i = 0; i < nums.length; i++) {
        let complement = target - nums[i];
        if (valToIdx.has(complement)) {
            return [valToIdx.get(complement), i];
        }
        valToIdx.set(nums[i], i);
    }
    return [];
};