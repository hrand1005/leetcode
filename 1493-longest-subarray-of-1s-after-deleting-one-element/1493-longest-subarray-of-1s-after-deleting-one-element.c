int longestSubarray(int* nums, int nums_size) {
    size_t l, r, zc;
    l = r = zc = 0;
    while (r < nums_size) {
        zc = (nums[r++]) ? zc : zc + 1;
        if (zc > 1 && nums[l++] == 0)
            zc--;
    }
    return r - l - 1;
}