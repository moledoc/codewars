// Create a function that takes a Roman numeral as its argument and returns its value as a numeric decimal integer.
// You don't need to validate the form of the Roman numeral.
package romannumeralsdecoder

// map roman numeral values to corresponding integer representation of the numeral.
var RuneValues = map[uint8]int{
	73: 1,    // I
	86: 5,    // V
	88: 10,   // X
	76: 50,   // L
	67: 100,  // C
	68: 500,  // D
	77: 1000, // M
}

func Decode(roman string) int {
	// init variable to save previous numeral value
	var prev int = 0
	// init result variable
	var result int = 0
	// from right to left value roman numerals and add them together.
	// if prev numeral value was greater than the current one,
	// then it means we need to subtract the smaller one from the result.
	// For example in case of IX (=9), XC (=90), CM (=900).
	for i := len(roman) - 1; i >= 0; i-- {
		var runeNumeral uint8 = roman[i]
		switch {
		case RuneValues[runeNumeral] < prev:
			result -= RuneValues[runeNumeral]
		case RuneValues[runeNumeral] >= prev:
			result += RuneValues[runeNumeral]
		}
		prev = RuneValues[runeNumeral]
	}
	return result
}
