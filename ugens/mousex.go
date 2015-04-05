package ugens

import . "github.com/briansorahan/sc/types"

// MouseX allpass delay with cubic interpolation
type MouseX struct {
	// Min is the value of this ugen's output when the
	// mouse is at the left edge of the screen
	Min Input
	// Max is the value of this ugen's output when the
	// mouse is at the right edge of the screen
	Max Input
	// Warp is the mapping curve. 0 is linear, 1 is exponential
	Warp Input
	// Lag factor to dezipper cursor movements
	Lag Input
}

func (self *MouseX) defaults() {
	if self.Min == nil {
		self.Min = C(0)
	}
	if self.Max == nil {
		self.Max = C(1)
	}
	if self.Warp == nil {
		self.Warp = C(0)
	}
	if self.Lag == nil {
		self.Lag = C(0.2)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (self MouseX) Rate(rate int8) Input {
	checkRate(rate)
	(&self).defaults()
	return UgenInput("MouseX", rate, 0, self.Min, self.Max, self.Warp, self.Lag)
}
