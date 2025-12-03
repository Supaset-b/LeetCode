func isPalindrome(x int) bool {

    if x < 0 {
        return false
    }

    originalX, reverseX := x, 0
    for x != 0 {
        reverseX = reverseX*10 + x%10
        x = x/10
    }

    return originalX == reverseX
}