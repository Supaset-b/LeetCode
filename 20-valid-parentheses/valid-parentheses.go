func isValid(s string) bool {
    stack := make([]rune, 10000)
    index := -1
    // for loop checking character in string.
    // if there is open bracket, push in to stack
    for _,r := range s {
        switch r {
            case '(','{','[':
                // push to on top stack.
                index ++
                stack[index] = r
            case ')':
                if index == -1 {
                    return false
                } else if stack[index] == '(' {
                    index--
                    continue
                }
                return false
            case '}':
                 if index == -1 {
                    return false
                } else if stack[index] == '{' {
                    index--
                    continue
                }
                return false
            case ']':
                 if index == -1 {
                    return false
                } else if stack[index] == '[' {
                    index--
                    continue
                }
                return false
            default:
                return false
        }
    }
    if index != -1 {
        return false
    }
    return true
}