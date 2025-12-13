func singleNumber(nums []int) int {
    // m := make(map[int]bool)
    // for _,num := range nums{
    //     if m[num] {
    //         m[num] = false
    //     } else {
    //         m[num] = true
    //     }
    // }
    // for k,v := range m {
    //     if v {
    //         return k
    //     }
    // }
    ret := 0
    for _,num := range nums {
        ret ^= num
    }
    return ret
}