int removeDuplicates(int* nums, int numsSize){
    int insert = 1;
    int last = nums[0];
    
    for (int i = 1; i < numsSize; i++) {
        if (nums[i] != last) {
            nums[insert] = nums[i];
            last = nums[i];
            insert++;
        }
    }
    
    return insert;
}