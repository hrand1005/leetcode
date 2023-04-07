bool prefix_equal(int i, char** strs, int strs_len) {
    char cmp = strs[0][i];
    
    for (int j = 0; j < strs_len; j++) {
        if (strlen(strs[j]) < i+1 || strs[j][i] != cmp) {
            return false;
        }
    }
    return true;
}

char * longestCommonPrefix(char ** strs, int strsSize){
    char* prefix;
    int i;
    
    if (strsSize == 0) {
        return false;
    }
    
    i = 0;
    while (prefix_equal(i, strs, strsSize)) {
        i++;
    }
    
    prefix = malloc(sizeof(char)*i+1);
    memset(prefix, '\0', i+1);
    strncpy(prefix, strs[0], i);
    return prefix;
}