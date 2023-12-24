package day20

import (
	"fmt"
	"strings"
)

type State int

const (
	StateOff State = iota
	StateOn
)

func (s State) String() string {
	return [...]string{"off", "on"}[s]
}

type PulseStrength int

const (
	PulseStrengthLow PulseStrength = iota
	PulseStrengthHigh
	PulseStrengthUndefined
)

func (ps PulseStrength) String() string {
	return [...]string{"low", "high", "undefined"}[ps]
}

type Pulse struct {
	Src      string
	Dst      string
	Strength PulseStrength
}

func (p Pulse) PrintPulse() {
	fmt.Printf("%s -%s-> %s\n", p.Src, p.Strength.String(), p.Dst)
}

type Pulser interface {
	Pulse(p *Pulse) []*Pulse
}

type Module struct {
	Label  string
	Inputs []string
	Output []string
}

func (m Module) sendPulses(s PulseStrength) []*Pulse {
	pulses := make([]*Pulse, 0)

	for _, o := range m.Output {
		pulses = append(pulses, &Pulse{
			Src:      m.Label,
			Dst:      o,
			Strength: s,
		})
	}

	return pulses
}

type ButtonModule struct {
	Module
}

func (b *ButtonModule) Pulse(p *Pulse) []*Pulse {
	return b.sendPulses(PulseStrengthLow)
}

type BroadcasterModule struct {
	Module
}

func (b *BroadcasterModule) Pulse(p *Pulse) []*Pulse {
	return b.sendPulses(p.Strength)
}

type FlipFlopModule struct {
	Module
	State State
}

func (f *FlipFlopModule) Pulse(p *Pulse) []*Pulse {
	if p.Strength == PulseStrengthHigh {
		return make([]*Pulse, 0)
	}

	var pulseToSend PulseStrength

	switch f.State {
	case StateOff:
		f.State = StateOn
		pulseToSend = PulseStrengthHigh
	case StateOn:
		f.State = StateOff
		pulseToSend = PulseStrengthLow
	}

	return f.sendPulses(pulseToSend)
}

type ConjunctionModule struct {
	Module
	prevPulses map[string]PulseStrength
}

func (c *ConjunctionModule) Pulse(p *Pulse) []*Pulse {
	pulseToSend := PulseStrengthLow

	c.prevPulses[p.Src] = p.Strength

	for _, v := range c.prevPulses {
		if v == PulseStrengthLow {
			pulseToSend = PulseStrengthHigh
			break
		}
	}

	return c.sendPulses(pulseToSend)
}

type OutputModule struct {
	Module
	prevPulse PulseStrength
}

func (o *OutputModule) Pulse(p *Pulse) []*Pulse {
	o.prevPulse = p.Strength
	return make([]*Pulse, 0)
}

type PulseQueue struct {
	pulses []*Pulse
}

func (pq *PulseQueue) Push(p *Pulse) int {
	pq.pulses = append(pq.pulses, p)

	return len(pq.pulses)
}

func (pq *PulseQueue) Pop() *Pulse {
	if pq.IsEmpty() {
		return nil
	}

	var pulse *Pulse

	pulse, pq.pulses = pq.pulses[0], pq.pulses[1:]

	return pulse
}

func (pq *PulseQueue) IsEmpty() bool {
	return pq.Size() == 0
}

func (pq *PulseQueue) Size() int {
	return len(pq.pulses)
}

type System struct {
	ButtonCount int
	HiPulses    int
	LowPulses   int
	Modules     map[string]Pulser
	Queue       PulseQueue
}

func (s *System) PushButton(showPulses bool) {
	s.ButtonCount++

	button := s.Modules["button"]

	buttonPulses := button.Pulse(nil)
	s.PushPulses(buttonPulses)

	for !s.Queue.IsEmpty() {
		next := s.ProcessPulse(showPulses)
		nextModule := s.Modules[next.Dst]

		var pulses []*Pulse
		if nextModule != nil {
			pulses = s.Modules[next.Dst].Pulse(next)
		}
		s.PushPulses(pulses)
	}
}

func (s *System) PushPulses(pulses []*Pulse) {
	for _, p := range pulses {
		switch p.Strength {
		case PulseStrengthLow:
			s.LowPulses++
		case PulseStrengthHigh:
			s.HiPulses++
		}
		s.Queue.Push(p)
	}
}

func (s *System) ProcessPulse(showPulse bool) *Pulse {
	p := s.Queue.Pop()

	if showPulse {
		p.PrintPulse()
	}

	return p
}

func (s *System) ResetPulseCounts() {
	s.HiPulses = 0
	s.LowPulses = 0
}

func (s *System) PrintSystemState() {
	fmt.Printf("Presses: %4d, Hi: %3d, Low: %3d, Total: %4d\n", s.ButtonCount, s.HiPulses, s.LowPulses, s.HiPulses+s.LowPulses)
}

func Part1(input []string, buttonPresses int, showPulses bool) int {
	s := parseInput(input)

	for i := 0; i < buttonPresses; i++ {
		s.PushButton(showPulses)
	}

	return s.HiPulses * s.LowPulses
}

func Part2(input []string, showPulses bool) int {
	s := parseInput(input)

	outputModule := s.Modules["rx"].(*OutputModule)
	for outputModule.prevPulse != PulseStrengthLow {
		s.PushButton(showPulses)
		s.ResetPulseCounts()
	}

	return s.ButtonCount
}

func parseInput(input []string) System {
	s := System{
		Modules: make(map[string]Pulser),
	}

	button := &ButtonModule{
		Module: Module{
			Label:  "button",
			Output: []string{"broadcaster"},
		},
	}

	s.Modules["button"] = button

	inputsMap := make(map[string][]string)
	for _, line := range input {
		split := strings.Split(line, "->")
		label := split[0]

		var inputLabel string
		switch rune(label[0]) {
		case '%':
			inputLabel = strings.Trim(label[1:], " ")
		case '&':
			inputLabel = strings.Trim(label[1:], " ")
		case 'b':
			inputLabel = strings.Trim(label, " ")
		}

		outputLabels := strings.Split(split[1], ", ")
		for _, l := range outputLabels {
			label := strings.Trim(l, " ")
			inputsMap[label] = append(inputsMap[label], inputLabel)
		}
	}

	for _, line := range input {
		split := strings.Split(line, "->")
		label := split[0]

		var inputLabel string
		switch rune(label[0]) {
		case '%':
			inputLabel = strings.Trim(label[1:], " ")
		case '&':
			inputLabel = strings.Trim(label[1:], " ")
		case 'b':
			inputLabel = strings.Trim(label, " ")
		}

		outputLabels := strings.Split(split[1], ", ")
		var outputs []string
		for _, l := range outputLabels {
			label := strings.Trim(l, " ")
			outputs = append(outputs, label)
		}

		switch rune(label[0]) {
		case '%':
			f := FlipFlopModule{
				Module: Module{
					Label:  inputLabel,
					Inputs: inputsMap[inputLabel],
					Output: outputs,
				},
				State: StateOff,
			}
			s.Modules[f.Label] = &f
		case '&':
			pulsesMap := make(map[string]PulseStrength)
			for _, i := range inputsMap[inputLabel] {
				pulsesMap[i] = PulseStrengthLow
			}
			c := ConjunctionModule{
				Module: Module{
					Label:  inputLabel,
					Inputs: inputsMap[inputLabel],
					Output: outputs,
				},
				prevPulses: pulsesMap,
			}
			s.Modules[c.Label] = &c
		case 'b':
			b := BroadcasterModule{
				Module: Module{
					Label:  inputLabel,
					Inputs: []string{"button"},
					Output: outputs,
				},
			}
			s.Modules[b.Label] = &b
		}
	}

	o := OutputModule{
		Module: Module{
			Label:  "rx",
			Inputs: []string{"zp"},
			Output: make([]string, 0),
		},
		prevPulse: PulseStrengthUndefined,
	}
	s.Modules["rx"] = &o

	return s
}
