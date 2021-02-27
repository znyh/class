package patterns

import (
	"testing"
)

func TestBridge(t *testing.T) {
	cpu := new(cpusoft)
	storage := new(storagesoft)

	p := new(phone)

	p.setbridge(cpu)
	p.work()

	p.setbridge(storage)
	p.work()
}
