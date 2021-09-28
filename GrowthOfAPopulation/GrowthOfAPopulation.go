package growthofapopulation

import (
	"math"
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
