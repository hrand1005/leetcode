int strStr(char * haystack, char * needle){
    int idx;
    
    char* start = strstr(haystack, needle);
    
    if (!start) {
        idx = -1;
    } else {
        int i = 0;
        char* phaystack = haystack;
        while (phaystack != start) {
            phaystack++;
            i++;
        }
        /*
        while (&haystack[i] != start) {
            i++;
        }
        */
        idx = i;
    }
    return idx;
}