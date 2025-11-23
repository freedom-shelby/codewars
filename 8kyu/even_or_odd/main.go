package even_or_odd

// IsEven returns true if the given number is even, false otherwise.
func IsEven(number int) bool {
	return (number % 2) == 0
}

// EvenOrOdd returns "Even" if the given number is even, "Odd" otherwise.
func EvenOrOdd(number int) string {
	if IsEven(number) {
		return "Even"
	}

	return "Odd"
}
