package atcodergo

func IsEven(i int) bool {
	return i%2 == 0
}

func ReverseString(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func Unique(slice []string) []string {
	m := make(map[string]bool)
	var uniqueSlice []string

	for _, v := range slice {
		if _, ok := m[v]; !ok {
			m[v] = true
			uniqueSlice = append(uniqueSlice, v)
		}
	}

	return uniqueSlice
}
