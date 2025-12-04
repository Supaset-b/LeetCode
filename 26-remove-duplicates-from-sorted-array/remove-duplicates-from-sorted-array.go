func removeDuplicates(nums []int) int {
    k := 0
    m :=  make(map[int]int)
    for _,num := range nums {
        if _, duplicate := m[num]; duplicate{
            continue
        }
        m[num] = num
        nums[k] = num
        k++
    }       
    return k
}