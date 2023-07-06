int compress(char* chars, int charsSize){
    char c;
    char buf[5];
    int count, idx, i, j;
    c = count = idx = 0;
    
    for (i = 0; i < charsSize; i++) {
        if (chars[i] == c)
            count++;
        else {
            if (count == 1)
                chars[idx++] = c;
            else if (count > 1) {
                chars[idx++] = c;
                memset(buf, 0, sizeof(buf));
                sprintf(buf, "%d", count);
                for (j = 0; j < strlen(buf); j++)
                    chars[idx++] = buf[j];
            }
            count = 1;
            c = chars[i];
        }
    }
    if (count == 1) {
        chars[idx++] = c;
    }
    else if (count > 1) {
        chars[idx++] = c;
        memset(buf, 0, sizeof(buf));
        sprintf(buf, "%d", count);
        for (i = 0; i < strlen(buf); i++)
            chars[idx++] = buf[i];
    }
    // chars[idx] = '\0';
    return idx;
}