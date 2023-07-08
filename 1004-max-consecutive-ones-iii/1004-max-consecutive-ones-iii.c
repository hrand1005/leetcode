int longestOnes(int* nums, int nums_size, int k){
    int l, r, max_seq, seq;
    l = r = max_seq = 0;
    
    while (r < nums_size) {
        if (nums[r++] == 0)
            k--;
        if (k < 0 && nums[l++] == 0)
            k++;
        seq = r - l;
        max_seq = (max_seq >= seq) ? max_seq : seq;
    }
    return max_seq;
}