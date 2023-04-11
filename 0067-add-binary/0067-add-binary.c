int max(int, int);

char * addBinary(char * a, char * b){
    int a_len;
    int b_len;
    
    int res_len;
    char* res;
    
    int aval;
    int bval;
    int carry;
    
    int dint;
    char dchar;
    
    a_len = strlen(a);
    b_len = strlen(b);
    
    res_len = max(a_len, b_len)+2; // +2 for extra digit, null terminator
    res = malloc(sizeof(char) * res_len);
    memset(res, '\0', res_len);
    
    carry = 0;
    for (int i = 0; i < res_len-1; i++) {
        aval = (a_len-1-i >= 0) ? (a[a_len-1-i] == '1') : 0;
        bval = (b_len-1-i >= 0) ? (b[b_len-1-i] == '1') : 0;
        dint = carry + aval + bval;
        dchar;
        switch(dint) {
            case 0:
                dchar = '0';
                carry = 0;
                break;
            case 1:
                dchar = '1';
                carry = 0;
                break;
            case 2:
                dchar = '0';
                carry = 1;
                break;
            case 3:
                dchar = '1';
                carry = 1;
                break;
        }
        res[res_len-2-i] = dchar;
    }
    
    if (res[0] == '0') {
        // shift contents left
        for (int i = 1; i < res_len-1; i++) {
            res[i-1] = res[i];
        }
        res[res_len-2] = '\0';
    }
    
    return res;
}
           
int max(int a, int b) {
    if (a > b) {
        return a;
    }
    return b;
}