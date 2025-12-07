func isPalindrome(s string) bool {

    var newS string
    for _, char := range s {
        if unicode.IsLetter(char) || unicode.IsDigit(char) {
            newS += string(char)
        }
    }
    newS = strings.ToLower(newS)
    
    i := 0
    j := len(newS)-1
    for i < j {
        if newS[i] != newS[j] {
            return false
        } 
        i ++
        j --
    }
    return true
}