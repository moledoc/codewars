package validbraces

import (
	"math"
)

// Mapping of braces pairs.
// The key is the opening brace.
// If we want to extend the solution to other braces (eg <> etc)
// or even define our own custom braces (eg ,;),
// then when we define them in this map, then this solution works for those cases as well.
var bracesPairs = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
}

func ValidBraces(str string) bool {
	// if there is odd number of characters in the string,
	// then one of them is not closed.
	if math.Mod(float64(len(str)), 2) != 0 {
		return false
	}

	// Make a map to hold opened braces (essencially used as a stack).
	var needClosing map[int]string = make(map[int]string, len(str))
	// Counter how many braces are currently opened.
	var opened int = 0

	// Loop over every elem in str.
	for _, elem := range str {
		var elemString string = string(elem)
		// check if it's opening brace (i.e. is elem the key in bracesPairs)
		_, isOpeningBrace := bracesPairs[elemString]
		// If `elem` is not opening brace,
		// then find the latest opened brace
		// and check if it can be closed with `elem`.
		// If not, then return false.
		// If yes, then delete the last element from the map (i.e. rm from stack)
		// and decrement counter `opened`.
		// Continue to the next `elem` in `str`.
		if !isOpeningBrace {
			var closeThis string = needClosing[len(needClosing)]
			if bracesPairs[closeThis] != elemString {
				return false
			}
			delete(needClosing, len(needClosing))
			opened -= 1
			continue
		}
		// If `elem` is opening brace, then increase the counter `opened`
		// and add the `elem` to the end of the map (i.e. stack).
		opened += 1
		needClosing[opened] = elemString

	}
	// In the end, the counter `opened` must be 0 (i.e. all braces are closed)
	return opened == 0
}
