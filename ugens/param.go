package ugens

import . "github.com/scgolang/sc/types"

// Params provides a way to add parameters to a synthdef
type Params struct {
	l []*Param
}

// Add adds a named parameter to a synthdef, with an initial value
func (self *Params) Add(name string, initialValue float32) Input {
	idx := len(self.l)
	p := newParam(name, int32(idx), initialValue)
	self.l = append(self.l, p)
	return p
}

// List returns a list of the params that have been
// added to a synthdef
func (self *Params) List() []*Param {
	return self.l
}

// Control returns a Ugen that should be used as
// the first ugen of any synthdef that has parameters
func (self *Params) Control() Ugen {
	return newControl(len(self.l))
}

// NewParams creates a new Params instance
func NewParams() *Params {
	p := Params{make([]*Param, 0)}
	return &p
}

type Param struct {
	name  string
	index int32
	val   float32
}

func (self *Param) Name() string {
	return self.name
}

func (self *Param) Index() int32 {
	return self.index
}

func (self *Param) GetInitialValue() float32 {
	return self.val
}

func (self *Param) Mul(in Input) Input {
	return BinOpMul(KR, self, in)
}

func (self *Param) Add(in Input) Input {
	return BinOpAdd(KR, self, in)
}

func (self *Param) MulAdd(mul, add Input) Input {
	return MulAdd(KR, self, mul, add)
}

func newParam(name string, index int32, initialValue float32) *Param {
	p := Param{name, index, initialValue}
	return &p
}

type Control struct {
	inputs  []Input
	outputs []Output
}

func (self *Control) Name() string {
	return "Control"
}

func (self *Control) Rate() int8 {
	return int8(1)
}

func (self *Control) SpecialIndex() int16 {
	return 0
}

func (self *Control) Inputs() []Input {
	return self.inputs
}

func (self *Control) Outputs() []Output {
	return self.outputs
}

func (self *Control) Add(val Input) Input {
	return self
}

func (self *Control) Mul(val Input) Input {
	return self
}

func (self *Control) MulAdd(mul, add Input) Input {
	return self
}

type ControlOutput struct{}

func (self *ControlOutput) Rate() int8 {
	return 1
}

func newControl(numOutputs int) Ugen {
	outputs := make([]Output, numOutputs)
	o := ControlOutput{}
	for i := 0; i < numOutputs; i++ {
		outputs[i] = &o
	}
	c := Control{make([]Input, 0), outputs}
	return &c
}
