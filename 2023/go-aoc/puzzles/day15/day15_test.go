package day15

import "testing"

var exInput = []string{"rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"}

func TestDay15(t *testing.T) {
	result := Day15(exInput)

	if result != 145 {
		t.Errorf("expected result to be 145, got=%v", result)
	}
}
