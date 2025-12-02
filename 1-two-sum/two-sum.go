func twoSum(nums []int, target int) []int {
    numsMap := make(map[int]int)
    for i,num := range nums {
        findSecondTarget := target-num
        if index, found := numsMap[findSecondTarget]; found {
            return []int{index, i} 
        }
        numsMap[num] = i
    }
    return nil
}
