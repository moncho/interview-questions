package mostFrequent

func MostFrequent(a []int) int {
	count := make(map[int]int)

	most := 0
	highestCount := 0
	for _, i := range a {
		count[i]++
		if count[i] > highestCount {
			highestCount = count[i]
			most = i
		}

	}

	return most
}
