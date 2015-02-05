package sc

import (
	"encoding/binary"
	"io"
)

// Variant
type Variant struct {
	Name               Pstring   `json:'name,omitempty'`
	InitialParamValues []float32 `json:'initialParamValues'`
}

// ReadVariant read a Variant from an io.Reader
func ReadVariant(r io.Reader, numParams int32) (*Variant, error) {
	name, err := ReadPstring(r)
	if err != nil {
		return nil, err
	}
	paramValues := make([]float32, numParams)
	for i := 0; int32(i) < numParams; i++ {
		err = binary.Read(r, byteOrder, &paramValues[i])
		if err != nil {
			return nil, err
		}
	}
	v := Variant{*name, paramValues}
	return &v, nil
}
