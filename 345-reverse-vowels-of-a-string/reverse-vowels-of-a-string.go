func reverseVowels(s string) string {
    left:=0
    right:=len(s)-1
    runeS := []rune(s)
    
    for left<=right {
        if isVowels(runeS[left]) && isVowels(runeS[right]) {
            // reverseVowels
            temp := runeS[left]
            runeS[left] = runeS[right]
            runeS[right] = temp
            left++
            right--
        } else if isVowels(runeS[left]) {
            right--
        } else if isVowels(runeS[right]) {
            left++
        } else {
            left++
            right--
        }
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