package main

type Pulse bool
type State bool

const (
	HIGH = true
	LOW  = false
	ON   = true
	OFF  = false
)

type Module struct{}

type FlipFlop struct {
	Module
	state  State
	output Module
}

func (f *FlipFlop) process(pulse Pulse) (Pulse, bool) {
	if pulse == HIGH {
		return LOW, false
	}
	f.state = !f.state
    return Pulse(f.state), true
}

type Conjuction struct {
    Module

}
