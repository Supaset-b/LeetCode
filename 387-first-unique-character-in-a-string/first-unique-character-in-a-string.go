func firstUniqChar(s string) int {

    mappingRuneIndex := make(map[rune]int)
    for _, c := range s {
        mappingRuneIndex[c]++
    }

    for i, c := range s {
        if mappingRuneIndex[c] == 1 {
            return i
        }
    }
    return -1
}