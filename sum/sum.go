package sum

// Sum adds numbers
func Sum(args ...int) (result int) {
	for _, a := range args {
		result += a
	}
	return
}
