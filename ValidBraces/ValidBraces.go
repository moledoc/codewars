// package main
package validbraces

import "fmt"

var bracesMapping = map[string]int{
	"(": 1,
	"[": 1,
	"{": 1,
	")": -1,
	"]": -1,
	"}": -1,
}

var parenthesisStatus = 0
var bracketStatus = 0
var curlyBracketStatus = 0

func ValidBraces(str string) bool {
	for _, char := range str {
		var charS string = string(char)
		switch {
		case charS == "(" || charS == ")":
			parenthesisStatus += bracesMapping[string(char)]
			if parenthesisStatus < 0 {
				return false
			}
		case charS == "[" || charS == "]":
			bracketStatus += bracesMapping[string(char)]
			if bracketStatus < 0 {
				return false
			}
		case charS == "{" || charS == "}":
			curlyBracketStatus += bracesMapping[string(char)]
			if curlyBracketStatus < 0 {
				return false
			}
		default:
			fmt.Println("Unknown brace type")
			return false
		}
	}

	return true
}

// func main() {
// 	fmt.Println(ValidBraces(")"))
// }
