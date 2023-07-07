bool isSubsequence(char * s, char * t){
    if (*s == '\0') 
        return true;
    
    for (; *t != '\0'; t++) {
        if (*s == *t)
            s++;
        if (*s == '\0')
            return true;
    }
    return false;
}