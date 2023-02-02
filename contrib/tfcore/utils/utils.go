package utils

// MustCombineMaps attempts to combine two maps. Panics if maps can not be combined.
func MustCombineMaps[T interface{}](m1, m2 map[string]T) map[string]T {
	for key, value := range m2 {
		_, exists := m1[key]
		if exists {
			panic("Key overlap found when combining maps")
		}
		m1[key] = value
	}
	return m1
}
