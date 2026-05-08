func numJewelsInStones(jewels string, stones string) int {
    m := make(map[rune]bool)
    for _,j := range jewels {
        m[j] = true
    }

    cnt := 0
    for _, s := range stones {
        if m[s] {
            cnt ++
        }
    }
    return cnt 
}