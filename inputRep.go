package sc

import (
	"encoding/binary"
	"io"
)

type InputRep struct {
	UgenIndex   int32 `json:"ugenIndex"`
	OutputIndex int32 `json:"outputIndex"`
}

// Write writes an inputSpec to an io.Writer
func (self *InputRep) Write(w io.Writer) error {
	if we := binary.Write(w, byteOrder, self.UgenIndex); we != nil {
		return we
	}
	return binary.Write(w, byteOrder, self.OutputIndex)
}

func readInputRep(r io.Reader) (*InputRep, error) {
	var ugenIndex, outputIndex int32
	err := binary.Read(r, byteOrder, &ugenIndex)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, byteOrder, &outputIndex)
	if err != nil {
		return nil, err
	}
	is := InputRep{ugenIndex, outputIndex}
	return &is, nil
}
