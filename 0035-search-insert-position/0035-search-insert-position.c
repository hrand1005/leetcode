int searchInsert(int* nums, int numsSize, int target){
    int l = 0;
    int r = numsSize-1;
    int mid = (r + l) / 2;
    
    while (l <= r) {
        if (nums[mid] == target) {
            return mid;
        } else if (nums[mid] < target) {
            l = mid+1;
        } else {
            r = mid-1;
        }
        mid = (r + l) / 2;
    }
    
    return (nums[mid] < target) ? mid + 1 : mid;
}