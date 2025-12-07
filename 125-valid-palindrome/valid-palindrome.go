func isPalindrome(s string) bool {

    lowerS := strings.ToLower(s)
    newS := ""
    for _, char := range lowerS {
        if unicode.IsLetter(char) || unicode.IsDigit(char) {
            newS += string(char)
        }
    }
    
    indexNewS := len(newS)-1
    for i,j := range newS {
        if j != rune(newS[indexNewS-i]) {
            return false
        }
    }
    return true
}
