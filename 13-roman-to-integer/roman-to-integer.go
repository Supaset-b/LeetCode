func romanToInt(s string) int {
    romanMap := make(map[string]int)
    romanMap["I"] = 1
    romanMap["V"] = 5
    romanMap["X"] = 10
    romanMap["L"] = 50
    romanMap["C"] = 100
    romanMap["D"] = 500
    romanMap["M"] = 1000

    numIndex := make([]int, len(s))
    result := 0
    for index, num := range s {
        strNum := string(num)
        numIndex[index] = romanMap[strNum]
        if index > 0 && numIndex[index] > numIndex[index-1] {
            result -= numIndex[index-1]*2
        }
        result += romanMap[strNum]
    }
    return result
}