bool isPalindrome(int x){
    char x_str[50];
    int n = sprintf(x_str, "%d", x);
    //printf("%d bytes written to x_str.\nx_str: %s\n", n, x_str);
    
    int x_str_len = strlen(x_str);
    //printf("length of x_str: %d\n", x_str_len);
    
    for (int i = 0; i < x_str_len/2; i++) {
        if (x_str[i] != x_str[x_str_len-1-i]) {
            return false;
        }
    }

    return true;
}