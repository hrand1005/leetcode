/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* productExceptSelf(int* nums, int numsSize, int* returnSize){
    int product, i;
    int forwards[numsSize];
    int backwards[numsSize];
    int *result = malloc(sizeof(int) * numsSize);
    
    forwards[0] = nums[0];
    for (i = 1; i < numsSize; i++) {
        forwards[i] = forwards[i-1] * nums[i];
    }
    
    backwards[numsSize-1] = nums[numsSize-1];
    for (i = numsSize-2; i >= 0; i--) {
        backwards[i] = backwards[i+1] * nums[i];
    }
    
    result[0] = backwards[1];
    result[numsSize-1] = forwards[numsSize-2];
    for (i = 1; i < numsSize-1; i++)
        result[i] = forwards[i-1] * backwards[i+1];
    
    *returnSize = numsSize;
    return result;
}