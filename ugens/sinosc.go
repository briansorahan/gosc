package ugens

import . "github.com/briansorahan/sc/types"

type SinOsc struct {
	Freq, Phase Input
}

func (self *SinOsc) defaults() {
	if self.Freq == nil {
		self.Freq = C(440)
	}
	if self.Phase == nil {
		self.Phase = C(0)
	}
}

func (self SinOsc) Rate(rate int8) *BaseNode {
	checkRate(rate)
	(&self).defaults()
	n := newNode("SinOsc", rate, 0)
	n.addInputs(self.Freq, self.Phase)
	return n
}
