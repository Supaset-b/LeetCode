func myAtoi(s string) int {
    checkSigned := true
    stringInt := ""
    for i := 0; i < len(s); i ++ {
        if (string(s[i]) >= "0" && string(s[i]) <= "9") || (string(s[i]) == "-" && checkSigned) || (string(s[i]) == "+" && checkSigned) {
            stringInt += string(s[i])
            checkSigned = false
        } else if string(s[i]) == " " {
            if len(stringInt) > 0 {
                break
            }
            continue
        } else {
            break
        }
    }
    
    result,_ := strconv.Atoi(stringInt)
    if result <= -2147483648 {
        return -2147483648
    } else if result >= 2147483647 {
        return 2147483647
    }
    return result
}