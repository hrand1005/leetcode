int convert_numeral(char c) {
    switch(c) {
        case 'I':
            return 1;
        case 'V':
            return 5;
        case 'X':
            return 10;
        case 'L':
            return 50;
        case 'C':
            return 100;
        case 'D':
            return 500;
        case 'M':
            return 1000;
    }
    return -1;
}

int romanToInt(char * s){
    int s_len = strlen(s);
    int total = 0;
    int last = convert_numeral(s[0]);
    
    for (int i = 0; i < s_len; i++) {
        int val = convert_numeral(s[i]);
        if (val > last) {
            total -= 2 * last;
        } 
        total += val;
        last = val;
    }
    
    return total;
}