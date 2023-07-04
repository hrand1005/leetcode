int max(int a, int b) {
    return (a > b) ? a : b;
}

char * mergeAlternately(char * word1, char * word2){
    size_t w1_len = strlen(word1);
    size_t w2_len = strlen(word2);
    size_t max_len = max(w1_len, w2_len);
    size_t result_len = w1_len + w2_len + 1;
    size_t idx = 0;
    
    char *result = malloc(sizeof(char) * (result_len));
    for (int i = 0; i < max_len; i++) {
        if (i < w1_len)
            result[idx++] = word1[i];
        if (i < w2_len)
            result[idx++] = word2[i];
    }
    result[result_len-1] = 0;
    return result;
}
