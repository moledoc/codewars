package ludicnumbers_test

import (
	"testing"

	lud "github.com/moledoc/codewars/LudicNumbers"
)

func TestLudicSum(t *testing.T) {
	if result := lud.SumLudic(1); result != 1 {
		t.Fatalf("Expected %v, got %v\n", 1, result)
	}
	if result := lud.SumLudic(8); result != 59 {
		t.Fatalf("Expected %v, got %v\n", 59, result)
	}
	if result := lud.SumLudic(10); result != 107 {
		t.Fatalf("Expected %v, got %v\n", 107, result)
	}
	if result := lud.SumLudic(25); result != 1100 {
		t.Fatalf("Expected %v, got %v\n", 1100, result)
	}
}
