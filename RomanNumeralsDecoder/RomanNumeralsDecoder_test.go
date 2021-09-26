package romannumeralsdecoder_test

import (
	"testing"

	decoder "github.com/moledoc/codewars/RomanNumeralsdeCoder"
)

type fn func(string) int

func singleArgTester(t *testing.T, fun fn, funArg string, expected int) {
	result := fun(funArg)
	if result != expected {
		t.Fatalf("Expected %v, got %v", expected, result)
	}
}

func TestRomanNumeralDecoder(t *testing.T) {
	singleArgTester(t, decoder.Decode, "XXI", 21)
	singleArgTester(t, decoder.Decode, "I", 1)
	singleArgTester(t, decoder.Decode, "XCIX", 99)
	singleArgTester(t, decoder.Decode, "IV", 4)
	singleArgTester(t, decoder.Decode, "MMVIII", 2008)
	singleArgTester(t, decoder.Decode, "MDCLXVI", 1666)
	singleArgTester(t, decoder.Decode, "MMXXI", 2021)
	singleArgTester(t, decoder.Decode, "MCMXCIX", 1999)
}
