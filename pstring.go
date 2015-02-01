package gosc

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

// Pstring is a pascal-format string, which is a byte containing
// the string length, followed by the bytes of the string
type Pstring struct {
	Length int8
	String string
}

func (self *Pstring) Write(w io.Writer) error {
	e := binary.Write(w, byteOrder, self.Length)
	if e != nil {
		return e
	}
	_, e = w.Write(bytes.NewBufferString(self.String).Bytes())
	return e
}

// NewPstring create a new Pstring
func NewPstring(s string) Pstring {
	length := len(s)
	return Pstring{int8(length), s}
}

// ReadPstring reads a Pstring from an io.Reader
func ReadPstring(r io.Reader) (*Pstring, error) {
	var length int8
	e := binary.Read(r, byteOrder, &length)
	if e != nil {
		return nil, e
	}
	s := make([]byte, length)
	read, e := r.Read(s)
	if e != nil {
		return nil, e
	}
	if read != int(length) {
		return nil, fmt.Errorf("could not read %d bytes", length)
	}
	ps := Pstring{
		length,
		bytes.NewBuffer(s).String(),
	}
	return &ps, nil
}
