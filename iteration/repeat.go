package iteration

// take a string to repeat and a number for number of time it should run
// and returns the string
func Repeat(character string, numOfIterations int) string {
	var repeated string
	for i := 0; i < numOfIterations; i++ {
		repeated += character
	}
	return repeated
}
