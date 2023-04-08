int removeDuplicates(int* nums, int numsSize){
    int insert = 1;
    int last = nums[0];
    
    for (int j = 1; j < numsSize; j++) {
        if (nums[j] != last) {
            nums[insert] = nums[j];
            last = nums[j];
            insert++;
        }
    }
    
    return insert;
}