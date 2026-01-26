func lengthOfLongestSubstring(s string) int {
    
    // check length first
    if len(s) == 0 {
        return 0
    }

    // using for loop, when find same character just going back to index after the character.
    cnt, startPosition := 0,0
    m := make(map[byte]int)
    for i:=0; i<len(s); i++ {
        // if already exist character.
        if m[s[i]] > startPosition {
            startPosition = m[s[i]]
        }

        // cnt change from index character - start position
        tempCnt := i+1 - startPosition

        if tempCnt > cnt {
            // update ret value
            cnt = tempCnt
        }
        
        m[s[i]] = i+1
    }
    return cnt
}