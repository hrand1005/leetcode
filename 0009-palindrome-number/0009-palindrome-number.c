/*
bool isPalindrome(int x){
    char x_str[20];
    sprintf(x_str, "%d", x);
    int x_str_len = strlen(x_str);
    
    for (int i = 0; i < x_str_len/2; i++) {
        if (x_str[i] != x_str[x_str_len-1-i]) {
            return false;
        }
    }

    return true;
}
*/
bool isPalindrome(int x){
    if (x < 0) {
        return false;
    }
    
    long int reverse = 0;
    int xcopy = x;
    
    
    while (xcopy != 0) {
        reverse *= 10;
        reverse += xcopy % 10;
        xcopy /= 10;
    }
    
    return (reverse == x);
}