package godemo

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	a := 2
	b := 3
	msg, err := Calculate(&a, &b, "addition")
	if msg != 5 || err != nil {
		t.Fatalf("Calculate Addition failed")
	}
	mmsg, err1 := Calculate(&a, &b, "multiplication")
	if mmsg != 7 || err1 != nil {
		t.Fatalf("Calculate multiplication failed")
	}
}
