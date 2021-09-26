package bouncingballs_test

import (
	"testing"

	bballs "github.com/moledoc/BouncingBalls"


func TestBouncingBalls(t *testing.T){
	var result int = bbals.BouncingBalls(3,0.66,1.5) 
	if result != 3 {
		t.Fatalf("Expected %v, got %v\n",3,result)
	}

}