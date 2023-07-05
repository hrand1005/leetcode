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
    size_t n;
    char *result;
    size_t s1_len = strlen(str1);
    size_t s2_len = strlen(str2);
    size_t concat_len = s1_len + s2_len + 1;
    char s1s2[concat_len];
    char s2s1[concat_len];
    
    memset(s1s2, '\0', concat_len);
    memset(s2s1, '\0', concat_len);
    
    strncpy(s1s2, str1, s1_len);
    strncat(s1s2, str2, s2_len);
    
    strncpy(s2s1, str2, s2_len);
    strncat(s2s1, str1, s1_len);
    
    if (strncmp(s1s2, s2s1, s1_len + s2_len) != 0)
        return calloc(sizeof(char), 1);
    
    n = gcd(s1_len, s2_len);
    result = malloc(sizeof(char) * n + 1);
    result[n] = '\0';
    strncpy(result, str1, n);
    return result;
}