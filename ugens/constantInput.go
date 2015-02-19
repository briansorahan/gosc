package ugens

type constantInput float32

func (self constantInput) IsConstant() bool {
	return true
}

func (self constantInput) Value() float32 {
	return float32(self)
}
