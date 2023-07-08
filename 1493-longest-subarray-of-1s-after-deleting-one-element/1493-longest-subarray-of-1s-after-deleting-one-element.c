/*
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
*/

int longestSubarray(int* nums, int nums_size) {
    size_t before, after, seq, max_seq, zero_found;
    before = after = seq = max_seq = zero_found = 0;
    
    for (size_t i = 0; i < nums_size; i++) {
        if (nums[i] == 0) {
            zero_found = 1;
            seq = before + after;
            max_seq = (max_seq >= seq) ? max_seq : seq;
            before = after;
            after = 0;
        } else
            after++;
    }
    seq = before + after;
    max_seq = (max_seq >= seq) ? max_seq : seq;
    return (zero_found) ? max_seq : max_seq - 1;
}