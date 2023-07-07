bool isSubsequence(char * s, char * t){
    if (*s == '\0') 
        return true;
    
    while (*t != '\0') {
        if (*s == *t++)
            s++;
        if (*s == '\0')
            return true;
    }
    return false;
}