func rotate(nums []int, k int)  {
    lengthNums := len(nums)
    tempNums := make([]int, lengthNums)
    pointer := k
    for _,n := range nums {
        if pointer >= lengthNums {
            for pointer >= lengthNums {
                pointer -= lengthNums
            }
        } 
        tempNums[pointer] = n 
        pointer ++
    }
    copy(nums, tempNums)
}