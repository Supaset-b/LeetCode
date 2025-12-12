func maxProfit(prices []int) int {
    min := prices[0]
    max := prices[0]
    result := 0
    for _, price := range prices {
        if price < min {
            min = price
            max = price
        } else if price > min && price > max {
            max = price
            if max-min > result {
                result = max - min
            }
        }
    }
    return result
}
