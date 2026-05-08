func fizzBuzz(n int) []string {
    
    stringN := make([]string, n)
    for i:=0;i<n;i++{
        if (i+1)%3==0 && (i+1)%5==0 {
            stringN[i] = "FizzBuzz"
        } else if (i+1)%3==0 {
            stringN[i] = "Fizz"
        } else if (i+1)%5==0 {
            stringN[i] = "Buzz"
        } else {
            stringN[i] = strconv.Itoa(i+1)
        }
    }
    return stringN
}