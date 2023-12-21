package day20

type State int

const (
	StateOff State = iota
	StateOn
)

type PulseStrength int

const (
	PulseUnspecified PulseStrength = iota
	PulseLow
	PulseHigh
)

type PulseProcessor interface {
	ProcesPulse(p Pulse) *Pulse
}

type Pulse struct {
	Strength PulseStrength
}

type FlipFlopModule struct { // prefix %
	State State
}

type ConjunctionModule struct { // prefix &
}

type ButtonModule struct{}

type BrodcastModule struct{}

type System struct {
	PulseCount int
}

func Part1() {
}
