int min(int a, int b) {
    return (a < b) ? a : b;
}

int max(int a, int b) {
    return (a > b) ? a : b;
}

int maxArea(int* height, int heightSize){
    int left, right, max_vol, vol;
    left = 0;
    right = heightSize - 1;
    max_vol = 0;
    
    while (left < right) {
        vol = (right - left) * min(height[left], height[right]);
        max_vol = max(max_vol, vol);
        
        if (height[left] < height[right])
            left++;
        else if (height[right] < height[left])
            right--;
        else {
            left++;
            right--;
        }
    }
    return max_vol;
}