func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    prefix := string(strs[0])
    preMaxlength := len(prefix)
    for _,str := range strs {
        tempPrefix := ""
        maxLength := 0
        for index, char := range str {
            if len(prefix) == 0 {
                return ""
            } else if preMaxlength < index || string(prefix[index]) != string(char) {
                break
            } else {
                tempPrefix += string(char)
                maxLength = index
            }
        }
       prefix = tempPrefix
       preMaxlength = maxLength
    }
    return prefix
}