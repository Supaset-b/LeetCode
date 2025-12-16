func plusOne(digits []int) []int {
    lengthDigit := len(digits)
    digits[lengthDigit-1] +=1
    for i:=lengthDigit-1; i>=0; i--{
        if digits[i]/10 > 0 {
            digits[i] = digits[i]%10
            if i == 0 {
                return append([]int{1}, digits...)
            }
            digits[i-1] += 1
        }
    }
    return digits
}
