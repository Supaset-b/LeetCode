func selfDividingNumbers(left int, right int) []int {
    var selfDividingNums []int
    
    for left<=right{
        if isSelfDividingNums(left) {
            selfDividingNums = append(selfDividingNums, left)
        } 
        left++
    }
    return selfDividingNums
}

func isSelfDividingNums(num int) bool {
    strNum := strconv.Itoa(num)
    
    for _,n := range strNum{
        if n == '0' {
            return false
        }
        val := int(n-'0')
        if num % val != 0 {
            return false
        }
    }
    return true
}