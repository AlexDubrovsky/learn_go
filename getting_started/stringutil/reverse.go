package stringutil

func Reverse(s string) string {
	temp := []rune(s)
	for i, j := 0, len(temp)-1; i < len(temp)/2; i, j = i+1, j-1 {
		temp[i], temp[j] = temp[j], temp[i]
	}
	return string(temp)
}
