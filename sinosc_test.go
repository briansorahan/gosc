package sc

import (
	. "github.com/scgolang/sc/types"
	. "github.com/scgolang/sc/ugens"
	"testing"
)

func TestSinOsc(t *testing.T) {
	name := "SineTone"
	def := NewSynthdef(name, func(params Params) Ugen {
		bus, freq := C(0), C(440)
		sine := SinOsc{Freq: freq}.Rate(AR)
		return Out{bus, sine}.Rate(AR)
	})
	same, err := def.Compare(`{
		Out.ar(0, SinOsc.ar(440));
    }`)
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdef different from sclang version")
	}
}
