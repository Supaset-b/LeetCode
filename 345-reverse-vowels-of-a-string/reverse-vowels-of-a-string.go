func reverseVowels(s string) string {
    left := 0
    right := len(s)-1
    runeS := []rune(s)
    
    for left<=right {
        leftIsVowel := isVowels(runeS[left])
        rightIsVewel := isVowels(runeS[right])
        if  leftIsVowel && rightIsVewel {
            temp := runeS[left]
            runeS[left] = runeS[right]
            runeS[right] = temp
            left++
            right--
        } else if leftIsVowel {
            right--
        } else if rightIsVewel {
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