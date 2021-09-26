package validbraces_test

import (
	"testing"

	validbraces "github.com/moledoc/codewars/ValidBraces"
)

func TestValidBraces(t *testing.T) {
	var testCases = map[string]bool{
		"(){}[]":             true,
		"((([[{[([])]}]])))": true,
		"([)]":               false,
		"](){}[]":            false,
		"{()":                false,
		"{":                  false,
		"}":                  false,
		"{}":                 true,
		"({({({({({({[{}":    false,
		"[[[[]":              false,
	}
	for braceString, braceBool := range testCases {
		var result bool = validbraces.ValidBraces(braceString)
		if result != braceBool {
			t.Fatalf("For %v excpected %v, got %v\n", braceString, braceBool, result)
		}
	}
}
