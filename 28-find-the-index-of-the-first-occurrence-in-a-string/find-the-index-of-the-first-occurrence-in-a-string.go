func strStr(haystack string, needle string) int {
    iNeedle := 0
    lengthNeedle := len(needle)
    lengthHaystack := len(haystack)
    returnIndex := -1
    if lengthNeedle > lengthHaystack {
        return -1
    }
    for i := 0; i < lengthHaystack; i++ {
        if haystack[i] == needle[iNeedle] {
            iNeedle++
            if iNeedle == 1 {
                returnIndex = i
            } 
            if iNeedle == lengthNeedle {
                return returnIndex
            } 
        } else {
            iNeedle = 0
            if returnIndex != -1 {
                i = returnIndex
                returnIndex = -1
            }
        }
    }
    return -1
}