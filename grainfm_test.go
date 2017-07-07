package sc

import "testing"

func TestGrainFM(t *testing.T) {
	name := "GrainFMExample"
	def := NewSynthdef(name, func(params Params) Ugen {
		bus := C(0)
		src := GrainFM{}.Rate(AR)
		return Out{bus, src}.Rate(AR)
	})
	same, err := def.CompareToFile("testdata/GrainFMExample.scsyndef")
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdef different from sclang version")
	}
}

// FIXME

// func TestGrainFM(t *testing.T) {
// 	const defName = "GrainFMTest"

// 	def := NewSynthdef(defName, func(p Params) Ugen {
// 		var (
// 			gate = p.Add("gate", 1)
// 			amp  = p.Add("amp", 1)

// 			bus     = C(0)
// 			mouseY  = MouseY{Min: C(0), Max: C(400)}.Rate(KR)
// 			freqdev = WhiteNoise{}.Rate(KR).Mul(mouseY)
// 		)
// 		env := Env{
// 			Levels:      []Input{C(0), C(1), C(0)},
// 			Times:       []Input{C(1), C(1)},
// 			Curve:       "sine",
// 			ReleaseNode: C(1),
// 		}
// 		ampenv := EnvGen{
// 			Env:        env,
// 			Gate:       gate,
// 			LevelScale: amp,
// 			Done:       FreeEnclosing,
// 		}.Rate(KR)

// 		var (
// 			trig     = Impulse{Freq: C(10)}.Rate(KR)
// 			modIndex = LFNoise{Interpolation: NoiseLinear}.Rate(KR).MulAdd(C(5), C(5))
// 			pan      = MouseX{Min: C(-1), Max: C(1)}.Rate(KR)
// 		)
// 		sig := GrainFM{
// 			NumChannels: 2,
// 			Trigger:     trig,
// 			Dur:         C(0.1),
// 			CarFreq:     C(440).Add(freqdev),
// 			ModFreq:     C(200),
// 			ModIndex:    modIndex,
// 			Pan:         pan,
// 		}.Rate(AR)

// 		return Out{bus, sig.Mul(ampenv)}.Rate(AR)
// 	})
// 	compareAndWrite(t, defName, def)
// }
