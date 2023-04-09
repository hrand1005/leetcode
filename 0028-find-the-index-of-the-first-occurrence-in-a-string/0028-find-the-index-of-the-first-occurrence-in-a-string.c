int strStr(char * haystack, char * needle){
    int idx;
    
    char* start = strstr(haystack, needle);
    
    if (!start) {
        idx = -1;
    } else {
        int i = 0;
        while (&haystack[i] != start) {
            i++;
        }
        idx = i;
    }
    return idx;
}