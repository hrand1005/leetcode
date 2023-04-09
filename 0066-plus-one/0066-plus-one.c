/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* plusOne(int* digits, int digitsSize, int* returnSize){
    int carry = 1;
    int* sum = malloc(sizeof(int) * digitsSize);
    int* result;
    
    for (int i = digitsSize-1; i >= 0; i--) {
        int digit = digits[i] + carry;
        if (digit > 9) {
            carry = 1;
            digit -= 10;
        } else {
            carry = 0;
        }
        sum[i] = digit;
    }
    
    if (carry) {
        *returnSize = digitsSize + 1;
        result = malloc(sizeof(int) * (*returnSize));
        memcpy(result+1, sum, digitsSize * sizeof(int));
        result[0] = carry;
    } else {
        *returnSize = digitsSize;
        result = malloc(sizeof(int) * (*returnSize));
        memcpy(result, sum, digitsSize * sizeof(int));
    }
    
    free(sum);
    return result;
}