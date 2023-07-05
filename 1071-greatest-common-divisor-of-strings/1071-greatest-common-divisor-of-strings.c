int gcd(int a, int b) {
    int tmp;
    while (b != 0) {
        tmp = a % b;
        a = b;
        b = tmp;
    }
    return a;
}

char * gcdOfStrings(char * str1, char * str2){
    char *result;
    
    size_t n;
    size_t s1_len = strlen(str1);
    size_t s2_len = strlen(str2);
    size_t concat_len = s1_len + s2_len + 1;
    
    char *s1s2 = calloc(sizeof(char), concat_len);
    char *s2s1 = calloc(sizeof(char), concat_len);
    
    strncpy(s1s2, str1, s1_len);
    strncat(s1s2, str2, s2_len);
    
    strncpy(s2s1, str2, s2_len);
    strncat(s2s1, str1, s1_len);
    
    if (strncmp(s1s2, s2s1, s1_len + s2_len) != 0)
        result = calloc(sizeof(char), 1);
    else {
        n = gcd(s1_len, s2_len);
        result = calloc(sizeof(char), n+1);
        strncpy(result, str1, n);
    }
        
    free(s1s2);
    free(s2s1);
    return result;
}