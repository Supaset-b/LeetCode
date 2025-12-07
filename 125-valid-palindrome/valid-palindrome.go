func isPalindrome(s string) bool {

    var newS string
    for _, char := range s {
        if unicode.IsLetter(char) || unicode.IsDigit(char) {
            newS += string(char)
        }
    }
    // newS = strings.ToLower(newS)
    
    i := 0
    j := len(newS)-1
    for i < j {
        if toLower(newS[i]) != toLower(newS[j]) {
            return false
        } 
        i ++
        j --
    }
    return true
}

func toLower(b byte) byte {
    // fast ASCII lowercase
    if b >= 'A' && b <= 'Z' {
        return b + 32
    }
    return b
}
