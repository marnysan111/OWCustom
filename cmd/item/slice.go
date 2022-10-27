package item

func DeleteSlice(item []string, target string) []string {
	var result []string
	for _, elememt := range item {
		if elememt != target {
			result = append(result, elememt)
		}
	}
	return result
}
