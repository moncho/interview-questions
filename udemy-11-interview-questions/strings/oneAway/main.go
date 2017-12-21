package oneAway

import "math"

func isOneAway(a, b string) bool {
	diff := math.Abs(float64(len(a) - len(b)))
	if diff > 1 {
		return false
	} else if diff == 1 {
		return isOneAwayDifLength(a, b)
	}
	modified := false
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			if modified {
				return false
			}
			modified = true
		}

	}

	return true

}

func isOneAwayDifLength(a, b string) bool {
	larger := a
	smaller := b
	if len(b) > len(a) {
		larger = b
		smaller = a
	}

	modifications := 0
	for i := 0; i < len(smaller); {
		if smaller[i] == larger[i+modifications] {
			i++
		} else {
			modifications++
			if modifications > 1 {
				return false
			}
		}
	}

	return true
}
