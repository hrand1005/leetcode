int lengthOfLastWord(char * s){
    int last_end;
    int last_start;
    int i = strlen(s)-1;
    
    while (i > 0 && s[i] == ' ') 
        i--;
    
    last_end = i;
    
    while (i >= 0 && s[i] != ' ')
        i--;
    
    last_start = i + 1;
    
    return (last_end - last_start) + 1;
}