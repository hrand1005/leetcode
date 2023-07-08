double findMaxAverage(int* nums, int numsSize, int k){
    int total, max_total, lo, hi;
    
    total = 0;
    for (int i = 0; i < k; i++)
        total += nums[i];
    
    max_total = total;
    
    lo = 0;
    hi = k;
    while (hi < numsSize) {
        total = total - nums[lo++] + nums[hi++];
        max_total = (max_total >= total) ? max_total : total;
    }
    return (double)max_total / k;
}