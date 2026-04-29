func reverseVowels(s string) string {
    left:=0
    right:=len(s)-1
    runeS := []rune(s)
    
    for left<=right {
        if !isVowels(runeS[left]) {
            left++
            continue
        } else if !isVowels(runeS[right]) {
            right--
            continue
        } 
        temp := runeS[left]
        runeS[left] = runeS[right]
        runeS[right] = temp
        left++
        right--
    }
    return string(runeS)
}

func isVowels(r rune) bool {
    switch r {
        case 'a','e','i','o','u','A','E','I','O','U':
            return true
        default:
            return false
    }
}