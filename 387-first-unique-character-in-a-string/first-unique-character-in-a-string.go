func firstUniqChar(s string) int {

    mappingIndexRune := make(map[rune]int)
    mappingRuneBool := make(map[rune]bool)
    
    for i, c := range s {
        if _, value := mappingIndexRune[c]; value{
            mappingRuneBool[c] = true
        } else {
            mappingRuneBool[c] = false
        }
        mappingIndexRune[c] = i
    }

    for index, value := range s {
        if !mappingRuneBool[value] {
            return index
        }
    }
    
    return -1
}