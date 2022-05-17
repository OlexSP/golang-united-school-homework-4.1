package reverse_string

/*
ReverseString() - function, that reverses the input string
the string may have international symbols,emojis and consist of several lines
*/

func ReverseString(input string) (output string) {
	runeSlice := []rune(input) // converting string to []rune
	revSlice := make([]rune, len(runeSlice))
	for i, v := range runeSlice {
		revSlice[len(runeSlice)-1-i] = v
	}
	output = string(revSlice)
	return output
}
