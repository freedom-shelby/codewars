package even_or_odd

import "testing"

func TestEvenOrOdd(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{2, "Even"},
		{1, "Odd"},
		{0, "Even"},
		{7, "Odd"},
		{-2, "Even"},
		{-1, "Odd"},
		{100, "Even"},
		{999, "Odd"},
	}

	for _, test := range tests {
		result := EvenOrOdd(test.input)
		if result != test.expected {
			t.Errorf("EvenOrOdd(%d) = %s; expected %s", test.input, result, test.expected)
		}
	}
}
