package sc

import (
	"fmt"
	"testing"
)

func TestServerStatus(t *testing.T) {
	s := NewServer(NetAddr{"127.0.0.1", DefaultPort})
	if s == nil {
		t.Fatal(fmt.Errorf("NewServer returned nil"))
	}
	err := s.Start()
	if err != nil {
		t.Fatal(err)
	}
}
