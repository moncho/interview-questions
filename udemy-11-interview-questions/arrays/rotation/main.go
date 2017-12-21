package rotation

func IsRotation(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range b {
		if v == a[0] {
			//return reflect.DeepEqual(a, append(b[i:], b[:i]...))
			return isRotationFromIndex(a, b, i)
		}
	}

	return false
}
func isRotationFromIndex(a, b []int, index int) bool {
	lenA := len(a)
	var posB int
	for i, v := range a {
		posB = (i + index) % lenA
		if v != b[posB] {
			return false
		}
	}
	return true
}
