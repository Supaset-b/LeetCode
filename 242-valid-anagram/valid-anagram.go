func isAnagram(s string, t string) bool {

    if len(s) != len(t) {
        return false
    }

    mapS := make(map[rune]int)
    mapT := make(map[rune]int)

    for _,charS := range s {
        mapS[charS]++
    }
    for _,charT := range t {
        mapT[charT]++
    }
    
    for k,_ := range mapT {
        if mapT[k] <= mapS[k] {
            continue
        } else {
         return false
        }
    }
    return true
}