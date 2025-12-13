func containsDuplicate(nums []int) bool {
    numMap := make(map[int]bool)
    for _,num := range nums {
        if _,dup := numMap[num]; dup{
            return true
        }
        numMap[num] = false
    }
    return false
}