// package growthofapopulation
package main

import (
	"flag"
	"fmt"
	"math"
	"time"
)

// use math to solve the problem
// helper function, where all args are converted to floats
// Proof of this solution can be found in this repo under GrowthOfAPopulation directory (this solution directory)
func nbYearFloat(p0 float64, percent float64, aug float64, p float64) int {
	// Handle case, when growth rate is 1 (meaning the geom series that is related with aug does not converge)
	if percent == 0 {
		return int((p - p0) / aug)
	}
	r := 1 + percent/100
	c := (aug + p*r - p) / (aug + p0*r - p0)
	return int(math.Ceil(math.Log(c) / math.Log(r)))
}

// use math to solve the problem
func NbYear(p0 int, percent float64, aug int, p int) int {
	return nbYearFloat(float64(p0), percent, float64(aug), float64(p))
}

// using iterative method to solve the problem
func NbYearForce(p0 int, percent float64, aug int, p int) int {
	var pIt float64 = float64(p0)
	var counter int
	for pIt < float64(p) {
		counter += 1
		pIt = (1+percent/100)*pIt + float64(aug)
	}
	return counter
}

func main() {
	p0 := flag.Int("p0", 0, "P_0")
	p := flag.Float64("p", 0, "percent")
	aug := flag.Int("aug", 0, "augment")
	pn := flag.Int("pn", 0, "P_N")
	flag.Parse()
	if (*p0) == 0 && (*p) == 0 && (*aug) == 0 && (*pn) == 0 {
		(*p0) = 100
		(*p) = 0.001
		(*aug) = 10
		(*pn) = 1000000000000000000
	}
	i := 100
	for i < (*pn) {
		startCalc := time.Now()
		NbYear(*p0, *p, *aug, i)
		endCalc := time.Since(startCalc)
		startIter := time.Now()
		NbYearForce(*p0, *p, *aug, i)
		endIter := time.Since(startIter)
		fmt.Printf("%-20v%-10v%-10s\n", i, endCalc, endIter)
		i *= 10
	}
}
