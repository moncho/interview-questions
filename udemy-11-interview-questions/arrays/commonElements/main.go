package commonElements

func CommonElements(a, b []int) []int {
	result := make([]int, 0)

	for i, j := 0, 0; i < len(a) && j < len(b); {
		if a[i] == b[j] {
			result = append(result, a[i])
			i++
			j++
		} else if a[i] > b[j] {
			j++
		} else {
			i++
		}

	}
	return result
}
