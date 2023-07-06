bool increasingTriplet(int* nums, int numsSize){
    int i, low, mid;
    low = mid = INT_MAX;
    
    for (i = 0; i < numsSize; i++) {
        if (nums[i] <= low)
            low = nums[i];
        else if (nums[i] <= mid) 
            mid = nums[i];
        else
            return true;
    }
    return false;
}