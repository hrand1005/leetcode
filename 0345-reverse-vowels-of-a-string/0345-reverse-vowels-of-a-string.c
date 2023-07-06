char * reverseVowels(char * s){
    char tmp;
    const char *vowels = "aeiouAEIOU";
    int left = 0;
    int right = strlen(s)-1;
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