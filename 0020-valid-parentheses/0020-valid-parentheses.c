bool parens_match(char open, char close) {
    bool is_match = false;
    
    if (open == '(' && close == ')' ||
       open == '{' && close == '}' ||
       open == '[' && close == ']') {
        is_match = true;
    }
    return is_match;
}

bool isValid(char * s){
    bool is_valid = true;
    
    const char* open = "({[";
    const char* closed = ")}]";
    
    int idx = 0;
    int s_len = strlen(s);
    char* queue = malloc(sizeof(char) * (s_len+1));
    
    for (int i = 0; i < s_len; i++) {
        if (strchr(open, s[i]) != NULL) {
            queue[idx] = s[i];
            idx++;
        } else if (strchr(closed, s[i]) != NULL) {
            if (idx == 0 || !parens_match(queue[idx-1], s[i])) {
                is_valid = false;
                break;
            }
            idx--;
        }
    }
    
    if (idx != 0) {
        is_valid = false;
    }
                
    free(queue);
    return is_valid;
}