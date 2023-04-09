/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* plusOne(int* digits, int digitsSize, int* returnSize){
    int digit;
    int carry = 1;
    int* sum = malloc(sizeof(int) * (digitsSize+1));
    
    for (int i = digitsSize-1; i >= 0; i--) {
        digit = digits[i] + carry;
        if (digit > 9) {
            carry = 1;
            digit -= 10;
        } else {
            carry = 0;
        }
        sum[i+1] = digit;
    }
    
    if (carry) {
        *returnSize = digitsSize + 1;
        sum[0] = carry;
    } else {
        *returnSize = digitsSize;
        // shift all elements left
        for (int i = 1; i < digitsSize+1; i++) {
            sum[i-1] = sum[i];
        }
    }
    
    return sum;
}