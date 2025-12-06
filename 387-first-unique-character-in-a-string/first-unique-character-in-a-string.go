func firstUniqChar(s string) int {

    // mappingRuneIndex := make(map[rune]int)
    // mappingRuneBool := make(map[rune]bool)
    
    // for i, c := range s {
    //     if _, value := mappingRuneIndex[c]; value{
    //         mappingRuneBool[c] = true
    //     } else {
    //         mappingRuneBool[c] = false
    //     }
    //     mappingRuneIndex[c] = i
    // }

    // for index, value := range s {
    //     if !mappingRuneBool[value] {
    //         return index
    //     }
    // }

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