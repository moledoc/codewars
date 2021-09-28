package growthofapopulation_test

import (
	"testing"

	growth "github.com/moledoc/codewars/GrowthOfAPopulation"
)

type testCases struct {
	p0       int
	percent  float64
	aug      int
	p        int
	expected int
}

func TestNbYear(t *testing.T) {
	c1 := testCases{1500, 5, 100, 5000, 15}
	c2 := testCases{1500000, 2.5, 10000, 2000000, 10}
	c3 := testCases{1500000, 0.25, 1000, 2000000, 94}
	c4 := testCases{1500000, 0.25, -1000, 2000000, 151}
	c5 := testCases{1500000, 0, 10000, 2000000, 50}
	cases := []testCases{c1, c2, c3, c4, c5}
	var result int
	for _, cs := range cases {
		result = growth.NbYear(cs.p0, cs.percent, cs.aug, cs.p)
		if result != cs.expected {
			t.Fatalf("With %v for %v expected %v, got %v\n", cs, "math sol", cs.expected, result)
		}
		result = growth.NbYearForce(cs.p0, cs.percent, cs.aug, cs.p)
		if result != cs.expected {
			t.Fatalf("With %v for %v expected %v, got %v\n", cs, "iterative sol", cs.expected, result)
		}
	}
}
