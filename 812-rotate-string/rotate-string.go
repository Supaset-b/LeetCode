func rotateString(s string, goal string) bool {
    // check length of input first
    if len(s) != len(goal) {
        return false
    }

    // check goal contains in s
    return strings.Contains(s+s, goal)
}