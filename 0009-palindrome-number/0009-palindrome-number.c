bool isPalindrome(int x){
    char x_str[50];
    sprintf(x_str, "%d", x);
    int x_str_len = strlen(x_str);
    
    for (int i = 0; i < x_str_len/2; i++) {
        if (x_str[i] != x_str[x_str_len-1-i]) {
            return false;
        }
    }

    return true;
}