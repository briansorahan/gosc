package ugens

import . "github.com/scgolang/sc/types"

// BrownNoise generates noise whose spectrum falls off in power by 6 dB per octave
type BrownNoise struct{}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (self BrownNoise) Rate(rate int8) Input {
	checkRate(rate)
	return UgenInput("BrownNoise", rate, 0, 1)
}
