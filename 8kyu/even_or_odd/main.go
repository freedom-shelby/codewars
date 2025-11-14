package even_or_odd

func IsEven(number int) bool {
	return (number % 2) == 0
}

func EvenOrOdd(number int) string {
	if IsEven(number) {
		return "Even"
	}

	return "Odd"
}
