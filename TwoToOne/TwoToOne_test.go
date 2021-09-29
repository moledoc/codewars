package twotoone_test

import (
	"testing"

	twotoone "github.com/moledoc/codewars/TwoToOne"
)

type testCase struct {
	s1       string
	s2       string
	expected string
}

func TestTwoToOne(t *testing.T) {
	testCase1 := testCase{s1: "aretheyhere", s2: "yestheyarehere", expected: "aehrsty"}
	testCase2 := testCase{s1: "loopingisfunbutdangerous", s2: "lessdangerousthancoding", expected: "abcdefghilnoprstu"}
	testCase3 := testCase{s1: "xyaabbbccccdefww", s2: "xxxxyyyyabklmopq", expected: "abcdefklmopqwxy"}
	testCase4 := testCase{s1: "abcdefghijklmnopqrstuvwxyz", s2: "abcdefghijklmnopqrstuvwxyz", expected: "abcdefghijklmnopqrstuvwxyz"}

	cases := []testCase{testCase1, testCase2, testCase3, testCase4}
	var result string
	for i, testcase := range cases {
		result = twotoone.TwoToOne(testcase.s1, testcase.s2)
		if result != testcase.expected {
			t.Fatalf("For case %v expected %v, got %v\n", i+1, testcase.expected, result)
		}
	}
}
