func addDigits(num int) int {
    sum := (num%10)+(num/10)
    for sum/10 != 0 {
        sum = (sum%10)+(sum/10)
    }
    return sum
}