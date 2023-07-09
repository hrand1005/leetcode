int pivotIndex(int* nums, int nums_size){
    int total, acc, i;
    total = acc = 0;
    for (int i = 0; i < nums_size; i++)
        total += nums[i];
    
    for (int i = 0; i < nums_size; i++) {
        if (total - nums[i] - acc == acc)
            return i;
        acc += nums[i];
    }
    return -1;
}