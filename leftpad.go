package leftpad


func Leftpad(phrase []rune, padding uint, fill rune) []rune {
	var result []rune
	for i := uint(0); i < padding; i++ {
		result = append(result, fill)
	}
	result = append(result, phrase...)
	return result
}

func LeftpadSpace(phrase []rune, padding uint) []rune {
	return Leftpad(phrase, padding, ' ')
}

func RJust(phrase []rune, length uint) []rune {
	diff := length - uint(len(phrase))

	if diff > 0 {
		return LeftpadSpace(phrase, diff)
	}

	return phrase
}