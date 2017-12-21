package nonRepeatingChar

import "strings"

func NonRepeatingChar(s string) rune {
	count := make(map[rune]int)
	var found string
	for _, c := range s {
		count[c]++
		if count[c] == 1 {
			found = found + string(c)
		} else if count[c] == 2 {
			index := strings.IndexRune(found, c)
			found = found[:index] + found[index+1:]
		}
	}
	if len(found) > 0 {
		return rune(found[0])
	}
	return 0
}
