package iteration

const repeatCount = 5

// Repeat repeats a given string
func Repeat(input string) (repeated string) {
	for i := 0; i < repeatCount; i++ {
		repeated += input
	}
	return repeated
}