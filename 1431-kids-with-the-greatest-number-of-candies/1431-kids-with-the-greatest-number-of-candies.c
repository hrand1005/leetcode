/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
bool* kidsWithCandies(int* candies, int candiesSize, int extraCandies, int* returnSize){
    int max_candies = 0; 
    for (int i = 0; i < candiesSize; i++)
        max_candies = (max_candies > candies[i]) ? max_candies : candies[i];
    
    bool *result = malloc(sizeof(bool) * candiesSize);
    for (int i = 0; i < candiesSize; i++)
        result[i] = (candies[i] + extraCandies) >= max_candies;
    
    *returnSize = candiesSize;
    return result;
}