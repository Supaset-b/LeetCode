func moveZeroes(nums []int)  {
    tempNums := make([]int, len(nums))
    copy(tempNums, nums)
    iZero := len(nums)-1
    index := 0
    for _, n := range tempNums {
        if n == 0 {
            nums[iZero] = n
            iZero--
        } else {
            nums[index] = n
            index++
        }
    }
}