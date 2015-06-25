package sc

import (
	. "github.com/scgolang/sc/types"
)

type UgenFunc func(p Params) Ugen

type PlayFunc func() Ugen

// Play corresponds to http://doc.sccode.org/Classes/Function.html#-play.
// This is syntactic sugar for
//     temp := DefaultServer.TempDefName()
//     def := NewSynthdef(temp, func(params Params) Ugen {
//     })
//     DefaultServer.SendDef(def)
//     sid := DefaultServer.NextSynthID()
//     DefaultServer.NewSynth(temp, sid, AddToHead, DefaultGroupID)
func Play(node Ugen) error {
	// To implement this we need a DefaultServer and
	// a way to generate the names of the temp synthdefs.
	// If the ugen node returned by f is not Out,
	// wrap the node in an Out ugen with bus 0.
	return nil
}

// Play(SinOsc.Ar(440))

// Once we support multichannel expansion...
// Play([]Ugen{SinOsc.Ar(440), SinOsc.Ar(441)})
