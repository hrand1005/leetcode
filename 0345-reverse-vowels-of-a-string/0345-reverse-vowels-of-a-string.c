char * reverseVowels(char * s){
    int left, right;
    char tmp;
    const char *vowels = "aeiouAEIOU";
    
    left = 0;
    right = strlen(s)-1;
    
    while (left < right) {
        while (left < right && !strchr(vowels, s[left]))
            left++;
        while (right > left && !strchr(vowels, s[right]))
            right--;
        
        tmp = s[left];
        s[left++] = s[right];
        s[right--] = tmp;
    }
    return s;
}