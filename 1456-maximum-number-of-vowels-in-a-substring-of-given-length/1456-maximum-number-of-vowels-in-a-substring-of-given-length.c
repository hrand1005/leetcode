int is_vowel(char v) {
    return strchr("aeiou", v) != NULL;
}

int maxVowels(char * s, int k){
    size_t vcount, maxv, l, r, slen;
    
    vcount = 0;
    for (size_t i = 0; i < k; i++)
        vcount += is_vowel(s[i]);
    
    maxv = vcount;
    l = 0;
    r = k;
    slen = strlen(s);
    while (r < slen) {
        if (maxv == k)
            return maxv;
        vcount = vcount - is_vowel(s[l++]) + is_vowel(s[r++]);
        maxv = (maxv >= vcount) ? maxv : vcount;
    }
    return maxv;
}