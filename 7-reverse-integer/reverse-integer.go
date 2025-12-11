func reverse(x int) int {

    isNegative := 1
    if x < 0{
        isNegative = -1
        x = x * -1
    }

    y := ""
    for x > 0 {
        y += strconv.Itoa(x%10)
        x = x/10
    }

    result,_ := strconv.Atoi(y)
    result = result * isNegative
    if result > 2147483647|| result < -2147483648 {
        return 0
    }
    return result
}