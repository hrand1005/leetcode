/**
 * @param {string} s
 * @return {number}
 */
var romanToInt = function(s) {
    const special = new Map([
        ["CM", 900], ["CD", 400],
        ["XC", 90], ["XL", 40],
        ["IX", 9], ["IV", 4],
    ]);
    const single = new Map([
        ["M", 1000], ["D", 500], ["C", 100],
        ["L", 50], ["X", 10], ["V", 5], ["I", 1]
    ]);
    
    let total = 0;
    while (s != "") {
        if (s.length > 1 && special.has(s.slice(0, 2))) {
            total += special.get(s.slice(0, 2));
            s = s.slice(2);
        } else {
            total += single.get(s[0]);
            s = s.slice(1);
        }
    }
    
    return total;
};