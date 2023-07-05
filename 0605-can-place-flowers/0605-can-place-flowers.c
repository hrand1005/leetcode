/*
bool canPlaceFlowers(int* flowerbed, int flowerbedSize, int n){
    size_t cap = 0;
    size_t pos = 0;
    while (pos < flowerbedSize) {
        bool before_valid = (pos == 0) || (flowerbed[pos-1] == 0);
        bool after_valid = (pos == flowerbedSize - 1) || (flowerbed[pos+1] == 0);
        if (before_valid && after_valid && (flowerbed[pos] == 0)) {
            cap++;
            pos += 2;
        } else
            pos++;
    }
    return (cap >= n);
}
*/

int placeable_flowers(int* flowerbed, int flowerbed_size) {
    if (flowerbed_size == 0) 
        return 0;
    if (flowerbed_size == 1)
        return flowerbed[0] == 0;
    int placeable = (flowerbed[0] == 0 && flowerbed[1] == 0);
    int increment = (placeable || flowerbed[0] == 1) ? 2 : 1;
    return placeable_flowers(&flowerbed[increment], flowerbed_size - increment) + placeable;
}

bool canPlaceFlowers(int* flowerbed, int flowerbedSize, int n){
    int max_placeable = placeable_flowers(flowerbed, flowerbedSize);
    return max_placeable >= n;
}
