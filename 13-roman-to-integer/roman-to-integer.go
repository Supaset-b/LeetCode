func romanToInt(s string) int {
    romanMap := map[string]int{"I":1, "V":5, "X":10, "L":50, "C":100, "D":500, "M":1000}
    // romanMap["I"] = 1
    // romanMap["V"] = 5
    // romanMap["X"] = 10
    // romanMap["L"] = 50
    // romanMap["C"] = 100
    // romanMap["D"] = 500
    // romanMap["M"] = 1000


    numIndex := make([]int, len(s))
    result := 0
    for index, num := range s {
        strNum := string(num)
        numIndex[index] = romanMap[strNum]
        if index > 0 && numIndex[index] > numIndex[index-1] {
            result -= numIndex[index-1]*2
        }
        result += numIndex[index]
    }
    return result
}