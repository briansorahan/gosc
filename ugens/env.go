package ugens

// 0, 3, -99, -99, -- starting level, num segments, releaseNode, loopNode
// 1, 0.1, 5, 4, -- first segment: level, time, curve type, curvature
// 0.5, 1, 5, -4, -- second segment: level, time, curve type, curvature
// 0, 0.2, 5, 4 -- and so on

const (
	CurveStep        = 0
	CurveLinear      = 1
	CurveExponential = 2
	CurveSine        = 3
	CurveWelch       = 4
	CurveCustom      = 5
	CurveSquared     = 6
	CurveCubed       = 7
)

// Env is not a ugen, but rather a way to generate
// Control arrays that get handed to EnvGen
var Env = newEnv()

type Envelope interface {
	// InputsArray provides EnvGen with the data it needs
	// to get a list of inputs
	InputsArray() []interface{}
}

type envelopeImpl struct {
	levels      []interface{}
	times       []interface{}
	curvetype   int
	curveature  interface{}
	releaseNode interface{}
	loopNode    interface{}
}

func (self *envelopeImpl) InputsArray() []interface{} {
	numSegments := len(self.levels)
	arr := make([]interface{}, 4*numSegments)
	arr[0] = self.levels[0]
	arr[1] = float32(numSegments - 1)
	arr[2] = self.releaseNode
	arr[3] = self.loopNode
	for i, t := range self.times {
		arr[(4*i)+4] = self.levels[i+1]
		arr[(4*i)+5] = t
		arr[(4*i)+6] = float32(self.curvetype)
		arr[(4*i)+7] = self.curveature
	}
	return arr
}

type env struct {
}

// Perc http://doc.sccode.org/Classes/Env.html#*perc
func (self *env) Perc(args ...interface{}) Envelope {
	defaults := []float32{0.01, 1, 1, -4}
	withDefaults := applyDefaults(defaults, args...)
	e := envelopeImpl{
		[]interface{}{float32(0), withDefaults[2], float32(0)},
		[]interface{}{withDefaults[0], withDefaults[1]},
		CurveCustom,
		withDefaults[3],
		float32(-99),
		float32(-99),
	}
	return &e
}

// Linen http://doc.sccode.org/Classes/Env.html#*linen
func (self *env) Linen(args ...interface{}) Envelope {
	defaults := []float32{0.01, 1, 1, 1, 1}
	withDefaults := applyDefaults(defaults, args...)
	levels := []interface{}{
		float32(0),
		withDefaults[3],
		withDefaults[3], // sustain level
		float32(0),
	}
	times := []interface{}{
		withDefaults[0],
		withDefaults[1],
		withDefaults[2],
	}
	var curve int
	switch v := withDefaults[4].(type) {
	case int:
		curve = v
	case float32:
		curve = int(v)
	default:
		panic("curve must be either int or float32")
	}
	e := envelopeImpl{
		levels,
		times,
		curve,
		float32(0),
		float32(-99),
		float32(-99),
	}
	return &e
}

func newEnv() *env {
	eg := env{}
	return &eg
}
