int strStr(char * haystack, char * needle){
    char* start = strstr(haystack, needle);
    
    if (!start) {
        return -1;
    } 
    
    int idx = 0;
    while (&haystack[idx] != start) 
        idx++;
    
    return idx;
}