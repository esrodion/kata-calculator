package RomanNumerals

import (
	"fmt"
	"strings"
)

var allowedSymbols map[int32]int

//// Exports

func init() {
	allowed := "IVXLCDM"
	values := [7]int{1, 5, 10, 50, 100, 500, 1000}

	allowedSymbols = make(map[int32]int)
	for i, v := range allowed {
		allowedSymbols[v] = values[i]
	}
}

func Encode(N int) (string, error) {
	if N > 3999 || N <= 0 {
		return "", fmt.Errorf("unable to encode integer %d as roman numeric: out of bounds [1, 3999]", N)
	}

	parts := make([]string, 4)

	parts[0] = convertDigit(N/1000, "", "", "M")
	parts[1] = convertDigit(N%1000/100, "M", "D", "C")
	parts[2] = convertDigit(N%100/10, "C", "L", "X")
	parts[3] = convertDigit(N%10, "X", "V", "I")

	return strings.Join(parts, ""), nil
}

func Decode(Str string) (int, error) {
	ok, err := stringValid(Str)
	if !ok {
		return 0, err
	}

	// go over string in reverse order
	runeArray := []rune(Str)
	sum := 0
	lastHigh := 0
	for i := len(runeArray) - 1; i >= 0; i-- {
		v := runeArray[i]
		val := allowedSymbols[v]
		switch {
		case val < lastHigh:
			sum -= val
		case val > lastHigh:
			lastHigh = val
			sum += val
		default:
			sum += val
		}
	}

	return sum, nil
}

//// Internal

func stringValid(Str string) (bool, error) {
	length := len(Str)

	// check length
	if length == 0 {
		return false, fmt.Errorf("unable to decode empty string as roman numeric")
	}

	// check for restricted symbols
	for _, s := range Str {
		if allowedSymbols[s] == 0 {
			return false, fmt.Errorf("unable to decode roman numeric %q: invalid symbols on line", Str)
		}
	}

	// two correct symbols
	if length == 2 {
		return true, nil
	}

	// check symbols order
	values := make([]int, length)
	runeArray := []rune(Str)
	for i := 0; i < length; i++ {
		values[i] = allowedSymbols[runeArray[i]]
	}

	for i := 2; i < length; i++ {
		if values[i] > values[i-1] && values[i] > values[i-2] {
			return false, fmt.Errorf("unable to decode roman numeric %q: invalid symbols order (%q%q%q)",
				Str, runeArray[i], runeArray[i-1], runeArray[i-2])
		} else if i >= 3 && values[i] == values[i-1] && values[i] == values[i-2] && values[i] == values[i-3] {
			return false, fmt.Errorf("unable to decode roman numeric %q: too many repeating symbols (%q%q%q%q)",
				Str, runeArray[i-3], runeArray[i-2], runeArray[i-1], runeArray[i])
		}
	}

	// string valid
	return true, nil
}

func convertDigit(Digit int, High string, Half string, Low string) string {
	switch {
	case Digit == 0:
		return ""
	case Digit <= 3:
		return strings.Repeat(Low, Digit)
	case Digit == 4:
		return Low + Half
	case Digit <= 8:
		return Half + strings.Repeat(Low, Digit-5)
	case Digit == 9:
		return Low + High
	default:
		return High // should never happen
	}
}
