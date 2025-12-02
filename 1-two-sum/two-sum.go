func twoSum(nums []int, target int) []int {
    numsMap := make(map[int]int)
    for i,num := range nums {
        if index, found := numsMap[target-num]; found {
            return []int{index, i} 
        }
        numsMap[num] = i
    }
    return nil
}
