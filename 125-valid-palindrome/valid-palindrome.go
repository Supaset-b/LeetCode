func isPalindrome(s string) bool {

    var newS string
    for index, char := range s {
        if isAlphaNumeric(s[index]) {
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

func isAlphaNumeric(c byte) bool{
    return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}