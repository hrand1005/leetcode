void moveZeroes(int* nums, int numsSize){
    int i, zidx, temp;
    zidx = 0;
    for (i = 0; i < numsSize; i++) {
        if (nums[i] != 0) {
            temp = nums[i];
            nums[i] = nums[zidx];
            nums[zidx++] = temp;
        }
    }
}