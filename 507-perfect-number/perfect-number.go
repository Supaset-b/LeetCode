func checkPerfectNumber(num int) bool {
    sumDividsors:=0
    for i:=num/2; i>0; i--{
        if num%i==0 {
            sumDividsors += i
        }
    }
    if sumDividsors!=num {
        return false
    }
    return true
}