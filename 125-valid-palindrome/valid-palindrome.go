func isPalindrome(s string) bool {

    newS := ""
    for _, char := range s {
        if unicode.IsLetter(char) || unicode.IsDigit(char) {
            newS += string(char)
        }
    }
    newS = strings.ToLower(newS)
    
    i, j := 0, len(newS)-1
    for i < j {
        if newS[i] != newS[j] {
            return false
        } 
        i ++
        j --
    }
    return true
}
