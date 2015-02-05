package sc

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

const (
	SYNTHDEF_START   = "SCgf"
	SYNTHDEF_VERSION = 2
)

var byteOrder = binary.BigEndian

type SynthDef struct {
	Name               string      `json,'name,omitempty'`
	NumConstants       int32       `json,'numConstants,omitempty'`
	Constants          []float32   `json,'constants,omitempty'`
	NumParams          int32       `json,'numParams,omitempty'`
	InitialParamValues []float32   `json,'initialParamValues,omitempty'`
	NumParamNames      int32       `json,'numParamNames,omitempty'`
	ParamNames         []ParamName `json,'paramNames,omitempty'`
	NumUgens           int32       `json,'numUgens,omitempty'`
	Ugens              []Ugen      `json,'ugens,omitempty'`
	NumVariants        int16       `json,'numVariants,omitempty'`
	Variants           []Variant   `json,'variants,omitempty'`
}

// Write writes a synthdef to an io.Writer
func (self *SynthDef) Write(w io.Writer) error {
	if he := self.writeHead(w); he != nil {
		return he
	}
	return self.writeBody(w)
}

// Dump writes information about a SynthDef to an io.Writer
func (self *SynthDef) Dump(w io.Writer) error {
	var e error

	fmt.Fprintf(w, "%-30s %s\n", "Name", self.Name)
	// write constants
	fmt.Fprintf(w, "%-30s %d\n", "NumConstants", self.NumConstants)
	fmt.Fprintf(w, "%s\n", "Constants")
	for i := 0; int32(i) < self.NumConstants; i++ {
		fmt.Fprintf(w, "    %-26d %g\n", i, self.Constants[i])
	}
	// write params
	fmt.Fprintf(w, "%-30s %d\n", "NumParams", self.NumParams)
	if self.NumParams > 0 {
		fmt.Fprintf(w, "%-30s\n", "Params:")
		for i := 0; int32(i) < self.NumParams; i++ {
			fmt.Fprintf(w, "    Initial Value %-12d %g\n", i, self.InitialParamValues[i])
		}
	}
	// write param names
	fmt.Fprintf(w, "%-30s %d\n", "NumParamNames", self.NumParamNames)
	if self.NumParamNames > 0 {
		fmt.Fprintf(w, "%-30s\n", "Param Names:")
		for i := 0; int32(i) < self.NumParamNames; i++ {
			fmt.Fprintf(w, "    %-26d %g\n", i, self.ParamNames[i])
		}
	}
	// write ugens and variants
	fmt.Fprintf(w, "%-30s %d\n", "NumUgens", self.NumUgens)
	fmt.Fprintf(w, "%-30s %d\n", "NumVariants", self.NumVariants)
	if self.NumUgens > 0 {
		for i := 0; int32(i) < self.NumUgens; i++ {
			fmt.Fprintf(w, "\nUgen %d:\n", i)
			e = self.Ugens[i].Dump(w)
			if e != nil {
				return e
			}
		}
	}
	if self.NumVariants > 0 {
		fmt.Fprintf(w, "%-30s\n", "Variants:")
		for i := 0; int16(i) < self.NumVariants; i++ {
			e = self.Ugens[i].Dump(w)
			if e != nil {
				return e
			}
		}
	}
	return nil
}

// write a synthdef header
func (self *SynthDef) writeHead(w io.Writer) error {
	_, we := w.Write(bytes.NewBufferString("SCgf").Bytes())
	if we != nil {
		return we
	}
	we = binary.Write(w, byteOrder, int32(SYNTHDEF_VERSION))
	if we != nil {
		return we
	}
	return binary.Write(w, byteOrder, int16(1))
}

// write a synthdef body
func (self *SynthDef) writeBody(w io.Writer) error {
	// write constants
	we := binary.Write(w, byteOrder, self.NumConstants)
	if we != nil {
		return we
	}
	for _, c := range self.Constants {
		if we = binary.Write(w, byteOrder, c); we != nil {
			return we
		}
	}
	// write parameters
	we = binary.Write(w, byteOrder, self.NumParams)
	if we != nil {
		return we
	}
	for _, p := range self.InitialParamValues {
		we = binary.Write(w, byteOrder, p)
		if we != nil {
			return we
		}
	}
	we = binary.Write(w, byteOrder, self.NumParamNames)
	if we != nil {
		return we
	}
	for _, p := range self.ParamNames {
		if we = p.Write(w); we != nil {
			return we
		}
	}
	// number of ugens
	if binary.Write(w, byteOrder, int32(1)); we != nil {
		return we
	}

	// TODO: write ugens

	// number of variants
	if we = binary.Write(w, byteOrder, int16(0)); we != nil {
		return we
	}
	return nil
}

func (self *SynthDef) Load(s Server) error {
	return nil
}

// ReadSynthDef reads a synthdef from an io.Reader
func ReadSynthDef(r io.Reader) (*SynthDef, error) {
	// read the type
	startLen := len(SYNTHDEF_START)
	start := make([]byte, startLen)
	read, er := r.Read(start)
	if er != nil {
		return nil, er
	}
	if read != startLen {
		return nil, fmt.Errorf("bad synthdef")
	}
	if bytes.NewBuffer(start).String() != SYNTHDEF_START {
		return nil, fmt.Errorf("bad synthdef")
	}
	// read version
	var version int32
	er = binary.Read(r, byteOrder, &version)
	if er != nil {
		return nil, er
	}
	if version != SYNTHDEF_VERSION {
		return nil, fmt.Errorf("bad synthdef version %d", version)
	}
	// read number of synth defs
	var numDefs int16
	er = binary.Read(r, byteOrder, &numDefs)
	if er != nil {
		return nil, er
	}
	if numDefs != 1 {
		return nil, fmt.Errorf("multiple synthdefs not supported")
	}
	// read synthdef name
	defName, er := ReadPstring(r)
	if er != nil {
		return nil, er
	}
	// read number of constants
	var numConstants int32
	er = binary.Read(r, byteOrder, &numConstants)
	if er != nil {
		return nil, er
	}
	// read constants
	constants := make([]float32, numConstants)
	for i := 0; i < int(numConstants); i++ {
		er = binary.Read(r, byteOrder, &constants[i])
		if er != nil {
			return nil, er
		}
	}
	// read number of parameters
	var numParams int32
	er = binary.Read(r, byteOrder, &numParams)
	if er != nil {
		return nil, er
	}
	// read initial parameter values
	initialValues := make([]float32, numParams)
	for i := 0; i < int(numParams); i++ {
		er = binary.Read(r, byteOrder, &initialValues[i])
		if er != nil {
			return nil, er
		}
	}
	// read number of parameter names
	var numParamNames int32
	er = binary.Read(r, byteOrder, &numParamNames)
	if er != nil {
		return nil, er
	}
	// read param names
	paramNames := make([]ParamName, numParamNames)
	for i := 0; int32(i) < numParamNames; i++ {
		pn, er := ReadParamName(r)
		if er != nil {
			return nil, er
		}
		paramNames[i] = *pn
	}
	// read number of ugens
	var numUgens int32
	er = binary.Read(r, byteOrder, &numUgens)
	if er != nil {
		return nil, er
	}
	// read ugens
	ugens := make([]Ugen, numUgens)
	for i := 0; int32(i) < numUgens; i++ {
		ugen, er := ReadUgen(r)
		if er != nil {
			return nil, er
		}
		ugens[i] = *ugen
	}
	// read number of variants
	var numVariants int16
	er = binary.Read(r, byteOrder, &numVariants)
	if er != nil {
		return nil, er
	}
	// read variants
	variants := make([]Variant, numVariants)
	for i := 0; int16(i) < numVariants; i++ {
		v, er := ReadVariant(r, numParams)
		if er != nil {
			return nil, er
		}
		variants[i] = *v
	}
	// return a new synthdef
	synthDef := SynthDef{
		defName.String(),
		numConstants,
		constants,
		numParams,
		initialValues,
		numParamNames,
		paramNames,
		numUgens,
		ugens,
		numVariants,
		variants,
	}
	return &synthDef, nil
}

// NewSynthDef creates a new SynthDef from a UgenGraphFunc
func NewSynthDef(name string, f UgenGraphFunc) *SynthDef {
	// this function has to be able to traverse a ugen
	// graph and turn it into a synth def
	return nil
}
