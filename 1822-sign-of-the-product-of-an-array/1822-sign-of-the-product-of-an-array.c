int arraySign(int* nums, int numsSize){
    int res = 1;
    for (int i = 0; i < numsSize; i++) {
        if (nums[i] == 0)
            return 0;
        else if (nums[i] < 0) 
            res *= -1;
    }
    return res;
}