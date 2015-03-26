package ugens

import (
	. "github.com/briansorahan/sc/types"
)

type C float32

func (self C) Val() float32 {
	return float32(self)
}

func (self C) Mul(val Input) Input {
	switch v := val.(type) {
	case *Node:
		return v.Mul(self)
	case C:
		return C(v * self)
	default:
		panic("input was neither ugen nor constant")
	}
}

func (self C) Add(val Input) Input {
	switch v := val.(type) {
	case *Node:
		return v.Add(self)
	case C:
		return C(v + self)
	default:
		panic("input was neither ugen nor constant")
	}
}

func (self C) Equals(val C) bool {
	return float32(self) == float32(val)
}
