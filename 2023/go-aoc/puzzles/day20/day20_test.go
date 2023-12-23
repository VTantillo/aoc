package day20

import (
	"fmt"
	"testing"
)

var config1 = []string{
	"broadcaster -> a, b, c",
	"%a -> b",
	"%b -> c",
	"%c -> inv",
	"&inv -> a",
}

var config2 = []string{
	"broadcaster -> a",
	"%a -> inv, con",
	"&inv -> b",
	"%b -> con",
	"&con -> output",
}

func TestPart1(t *testing.T) {
	result1 := Part1(config1, 1000)

	if result1 != 32000000 {
		t.Fatalf("Result should have been 32000000, got=%v ", result1)
	}

	result2 := Part1(config2, 1000)

	if result2 != 11687500 {
		t.Fatalf("Result should have been 11687500, got=%v ", result2)
	}
}

func TestPushButton(t *testing.T) {
	sys := parseInput(config1)

	sys.PushButton()

	if sys.LowPulses != 8 {
		t.Fatalf("Low pulses should have been 8, got=%v", sys.LowPulses)
	}

	if sys.HiPulses != 4 {
		t.Fatalf("Low pulses should have been 4, got=%v", sys.HiPulses)
	}
}

func TestModuleCreation(t *testing.T) {
	sys := parseInput(config2)

	for key, module := range sys.Modules {
		fmt.Printf("label: %s:  %T\n", key, module)
	}

	if len(sys.Modules) != 6 {
		t.Fatalf("Should have been 6 modules, got=%v", len(sys.Modules))
	}
}
