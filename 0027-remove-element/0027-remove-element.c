int removeElement(int* nums, int numsSize, int val){
    int idx = 0;
    
    for (int i = 0; i < numsSize; i++) {
        nums[idx] = nums[i];
        if (nums[idx] != val) {
            idx++;
        }
    }
    
    return idx;
}