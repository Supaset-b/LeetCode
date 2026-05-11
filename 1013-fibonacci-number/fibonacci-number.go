func fib(n int) int {
    if n == 0 {
        return 0
    }
    prev := 0
    curr := 1

    for i:=2; i<=n; i++ {
        prev, curr = curr, prev + curr
    }
    return curr
}
