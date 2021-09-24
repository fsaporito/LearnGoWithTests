package iteration

// Repeats the given character N times
func Repeat(char string, N int) string {
	if N <= 0 {
		return char
	}
	var repeated string
	for i := 0; i < N; i++ {
		repeated += char
	}
	return repeated
}
