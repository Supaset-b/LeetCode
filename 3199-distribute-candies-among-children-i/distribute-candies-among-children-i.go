func distributeCandies(n int, limit int) int {

    totalCandies1 := 0
    totalCandies2 := 0
    output := 0
		// children number 1
		for children1 := limit; children1 >= 0; children1-- {

			// total candies from children1
			totalCandies1 = n - children1

			if totalCandies1 == 0 {
				output++
				continue
			}

			// children number 2
			for children2 := limit; children2 >= 0; children2-- {

				// total candies from children2
				totalCandies2 = totalCandies1 - children2

				if totalCandies2 == 0 {
					output++
					continue
				}

				// children number 3
				for children3 := limit; children3 >= 0; children3-- {

					if children1+children2+children3 != n {
						continue
					}

					if totalCandies2 != 0 {
						output++
						totalCandies2--
					}
				}
			}
		}
    return output
}