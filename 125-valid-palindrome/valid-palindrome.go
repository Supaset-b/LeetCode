func isPalindrome(s string) bool {

    newS := ""
    for _, char := range s {
        if unicode.IsLetter(char) || unicode.IsDigit(char) {
            newS += string(char)
        }
    }
    newS = strings.ToLower(newS)
    indexNewS := len(newS)-1
    for i,j := range newS {
        if j != rune(newS[indexNewS-i]) {
            return false
        }
    }
    return true
}
