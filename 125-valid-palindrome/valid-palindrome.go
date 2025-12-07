func isPalindrome(s string) bool {

    lowerS := strings.ToLower(s)
    newS := ""
    for _, char := range lowerS {
        if unicode.IsLetter(char) || unicode.IsDigit(char) {
            newS += string(char)
        }
    }
    
    for i,j := range newS {
        if string(j) != string(newS[len(newS)-i-1]) {
            return false
        }
    }

    return true
}
