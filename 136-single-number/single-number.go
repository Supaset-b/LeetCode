func singleNumber(nums []int) int {
    m := make(map[int]bool)
    for _,num := range nums{
        if m[num] {
            m[num] = false
        } else {
            m[num] = true
        }
    }
    for k,v := range m {
        if v {
            return k
        }
    }
    return 0
}