package patterns

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	s1 := getSingleton()
	s2 := getSingleton()

	if s1 != s2 {
		t.Error("bad singleton")
	}
}
