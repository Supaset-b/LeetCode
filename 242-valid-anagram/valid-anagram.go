func isAnagram(s string, t string) bool {

    mapS := make(map[rune]int)
    mapT := make(map[rune]int)

    if len(s) != len(t) {
        return false
    }
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