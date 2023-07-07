void moveZeroes(int* nums, int numsSize){
    int zero_idx = 0;
    for (int i = 0; i < numsSize; i++) {
        if (nums[zero_idx] == 0) {
            nums[zero_idx] = nums[i];
            nums[i] = 0;
        } 
        if (nums[zero_idx] != 0)
            zero_idx++;
    }
}