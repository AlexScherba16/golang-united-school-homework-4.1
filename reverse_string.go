package reverse_string

func ReverseString(input string) (output string) {
	runes := []rune(input)
	reversed := make([]rune, len(runes))

	for i := len(runes) - 1; i >= 0; i-- {
		reversed[len(runes)-i-1] = runes[i]
	}
	return string(reversed)
}
