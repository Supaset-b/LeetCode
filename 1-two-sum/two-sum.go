func twoSum(nums []int, target int) []int {
    numsMap := make(map[int]int)
    for i,num := range nums {
        findSecondNum := target-num
        if index, found := numsMap[findSecondNum]; found {
            return []int{index, i} 
        }
        numsMap[num] = i
    }
    return nil
}
