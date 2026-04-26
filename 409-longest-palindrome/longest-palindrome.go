func longestPalindrome(s string) int {
    m := make(map[rune]int)
    
    // store character to map
    for _,r := range s {
        m[r] += 1
    }
    
    ret := 0
    for _,v := range m {
        ret += (v/2)*2
    }

    if len(s) > ret {
        ret ++
    }

    return ret
}