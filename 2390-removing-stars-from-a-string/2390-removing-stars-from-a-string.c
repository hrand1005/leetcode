char * removeStars(char * s){
    size_t idx = 0;
    size_t len = strlen(s);
    char buf[len+1];
    
    for (size_t i = 0; i < len; i++) {
        if (s[i] == '*')
            idx--;
        else
            buf[idx++] = s[i];
    }
    buf[idx] = '\0';
    return strdup(buf);
}