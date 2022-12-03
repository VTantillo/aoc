package day2

import "testing"

func TestRound1(t *testing.T) {
	score := calculateRound("A", "Y")
	if score != 8 {
		t.Error("incorrect result: expected 8, got", score)
	}
}

func TestRound2(t *testing.T) {
	score := calculateRound("B", "X")
	if score != 1 {
		t.Error("incorrect result: expected 1, got", score)
	}
}

func TestRound3(t *testing.T) {
	score := calculateRound("C", "Z")
	if score != 6 {
		t.Error("incorrect result: expected 6, got", score)
	}
}

func TestRound1Step2(t *testing.T) {
	score := calculateRound2("A", "Y")
	if score != 4 {
		t.Error("incorrect result: expected 4, got", score)
	}
}

func TestRound2Step2(t *testing.T) {
	score := calculateRound2("B", "X")
	if score != 1 {
		t.Error("incorrect result: expected 1, got", score)
	}
}

func TestRound3Step2(t *testing.T) {
	score := calculateRound2("C", "Z")
	if score != 7 {
		t.Error("incorrect result: expected 7, got", score)
	}
}
