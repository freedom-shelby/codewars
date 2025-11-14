package even_or_odd

func IsEvenOrOdd(number int) bool {
	return number%2 == 0
}

func EvenOrOdd(number int) string {
	if IsEvenOrOdd(number) {
		return "Even"
	}

	return "Odd"
}
