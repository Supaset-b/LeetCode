func myAtoi(s string) int {
    checkBool := true
    stringInt := ""
    for i := 0; i < len(s); i ++ {
        if (string(s[i]) >= "0" && string(s[i]) <= "9") || (string(s[i]) == "-" && checkBool) || (string(s[i]) == "+" && checkBool) {
            stringInt += string(s[i])
            checkBool = false
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