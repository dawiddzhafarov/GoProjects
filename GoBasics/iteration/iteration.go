package iteration

func Repeat(letter string, counter int) string {
	var word string
	for i := 0; i < counter; i++ {
		word += letter
	}
	return word
}
