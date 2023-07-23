package lib

func DeDuplicate(slice []string) []string {
	uniqueMap := make(map[string]bool)
	deduplicatedSlice := []string{}

	for _, item := range slice {
		if !uniqueMap[item] {
			deduplicatedSlice = append(deduplicatedSlice, item)
			uniqueMap[item] = true
		}
	}

	return deduplicatedSlice
}
