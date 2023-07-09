int largestAltitude(int* gain, int gain_size){
    int h = 0, hmax = 0;
    for (int i = 0; i < gain_size; i++) {
        h += gain[i];
        hmax = (hmax >= h) ? hmax : h;
    }
    return hmax;
}