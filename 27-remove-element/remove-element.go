func removeElement(nums []int, val int) int {
    ret := 0
    tempNum := make([]int, len(nums))
    for _,n := range nums {
        if n==val {
            continue
        } else {
            tempNum[ret] = n
            ret++
        }
    }
    copy(nums, tempNum)
    return ret
}