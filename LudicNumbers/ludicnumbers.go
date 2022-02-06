package ludicnumbers

import (
	"fmt"
)

type test int

func (result test) Assert(expected test) {
	if result != expected {
		panic(fmt.Sprintf("Excpeted value %v, but got %v", expected, result))
	}
}

func SumLudic(n int) int {
	if n == 1 {
		return 1
	}
	sum := 1
	var ludic []int
	var steps []int
	i := 2
	for len(ludic) < (n - 1) {
		add := true
		for j := 0; j < len(ludic); j++ {
			steps[j]++
			if steps[j] == ludic[j] {
				add = false
				steps[j] = 0
				break
			}
		}
		if add {
			ludic = append(ludic, i)
			steps = append(steps, 0)
			sum += i
		}
		i++
	}
	return sum
}

// func main() {
// 	test1 := SumLudic(1)
// 	test2 := SumLudic(10)
// 	test2 := SumLudic(10)
// 	test3 := SumLudic(25)
// 	fmt.Println(test1, test2, test3)
// 	test(test1).Assert(test(1))
// 	test(test2).Assert(test(107))
// 	test(test3).Assert(test(1100))
// }
