int removeDuplicates(int* nums, int numsSize){
    int i = 1;
    int last = nums[0];
    
    for (int j = 1; j < numsSize; j++) {
        if (nums[j] != last) {
            nums[i] = nums[j];
            last = nums[j];
            i++;
        }
    }
    
    return i;
}