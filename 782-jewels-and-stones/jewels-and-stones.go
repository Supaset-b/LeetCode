func numJewelsInStones(jewels string, stones string) int {
    m := make(map[rune]bool)
    for _,j := range jewels {
        m[j] = true
    }

    cnt := 0
    l := 0
    r := len(stones)-1
    for l<=r {
        if m[rune(stones[l])] {
            cnt++
            if l==r {
                break
            }
        }
        if m[rune(stones[r])] {
            cnt++
        }
        l++
        r--
    }
    return cnt 
}