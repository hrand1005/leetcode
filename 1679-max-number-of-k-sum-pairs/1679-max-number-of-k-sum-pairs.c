int partition(int *nums, int left, int right) {
    int p_idx, p, temp;
    p_idx = left;
    p = nums[left++];
    
    for (int i = left; i < right; i++) {
        if (nums[i] < p) {
            temp = nums[i];
            nums[i] = nums[left];
            nums[left++] = temp;
        }
    }
    nums[p_idx] = nums[left-1];
    nums[left-1] = p;
    return left-1;
}

void quicksort(int *nums, int left, int right) {
    int p;
    
    if (left >= right)
        return;
    
    p = partition(nums, left, right);
    quicksort(nums, left, p);
    quicksort(nums, p+1, right);
}


int compare(void const *a, void const *b) {
    return *(int *)a - *(int *)b;
}

int maxOperations(int* nums, int numsSize, int k){
    int left, right, lnum, rnum, ops;
    
    qsort(nums, numsSize, sizeof(int), compare);
    
    left = 0;
    right = numsSize - 1;
    ops = 0;
    
    while (left < right) {
        lnum = nums[left];
        rnum = nums[right];
        if (lnum + rnum == k) {
            left++;
            right--;
            ops++;
        } else if (lnum + rnum < k)
            left++;
        else
            right--;
    }
    
    return ops;
}