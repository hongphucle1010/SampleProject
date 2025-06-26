package utils

func GenerateNextID[T any](items []T, getID func(T) int) int {
	minID := 220000
	maxID := minID - 1
	for _, item := range items {
		id := getID(item)
		if id >= minID && id <= 229999 && id > maxID {
			maxID = id
		}
	}
	return maxID + 1
}
