/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
bool* kidsWithCandies(int* candies, int candiesSize, int extraCandies, int* returnSize){
    int maxc = 0; 
    for (int i = 0; i < candiesSize; i++)
        maxc = (maxc > candies[i]) ? maxc : candies[i];
    
    bool *result = malloc(sizeof(bool) * candiesSize);
    for (int i = 0; i < candiesSize; i++)
        result[i] = (candies[i] + extraCandies) >= maxc;
    
    *returnSize = candiesSize;
    return result;
}