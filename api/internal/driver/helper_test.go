package driver_test

func assertMapEquals[TKey comparable, TValue comparable](a, b map[TKey]TValue) bool {
	if len(a) != len(b) {
		return false
	}
	for k, valA := range a {
		valB := b[k]
		if valA != valB {
			return false
		}
	}
	return true
}
