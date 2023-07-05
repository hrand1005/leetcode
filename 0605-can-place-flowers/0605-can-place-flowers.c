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