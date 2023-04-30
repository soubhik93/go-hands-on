package iteration

const times = 5

func Repeat(character string) string {
	var result string

	for i := 0; i < times; i++ {
		result += character
	}

	return result
}
