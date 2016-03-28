package sc

import (
	"github.com/scgolang/osc"
)

const (
	freeSynthNodeAddress = "/n_free"
	setSynthNodeAddress  = "/n_set"
)

// Synth encapsulates a synth node.
type Synth struct {
	DefName string `json:"defName"`
	ID      int32  `json:"id"`
	client  *Client
}

// Get the value of a synth control.
func (s *Synth) Get(controlName string) (float32, error) {
	return 0, nil
}

// Set the value of a synth control.
func (s *Synth) Set(ctls map[string]float32) error {
	set, err := osc.NewMessage(setSynthNodeAddress)
	if err != nil {
		return err
	}
	if err := set.WriteInt32(s.ID); err != nil {
		return err
	}
	for name, value := range ctls {
		if err := set.WriteString(name); err != nil {
			return err
		}
		if err := set.WriteFloat32(value); err != nil {
			return err
		}
	}
	return s.client.oscConn.Send(set)
}

// newSynth creates a new synth structure.
func newSynth(client *Client, defName string, id int32) *Synth {
	return &Synth{
		DefName: defName,
		ID:      id,
		client:  client,
	}
}
