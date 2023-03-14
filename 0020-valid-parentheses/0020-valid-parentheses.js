/**
 * @param {string} s
 * @return {boolean}
 */
const closedToOpen = new Map([
    ["}", "{"], [")", "("], ["]", "["],
])

var isValid = function(s) {
    let openQueue = []
    
    for (let i = 0; i < s.length; i++) {
        if (!closedToOpen.has(s[i])) {
            openQueue.push(s[i])
        } else {
            if (openQueue.length === 0) {
                return false
            }
            if (openQueue[openQueue.length-1] !== closedToOpen.get(s[i])) {
                return false
            }
            openQueue.pop()
        }
    }
    
    return openQueue.length === 0
};