package memento

import (
	"fmt"
	"testing"
)

func TestCaretaker_RecoverOriginator(t *testing.T) {
	originator := new(Originator)
	originator.SetState("init state")

	caretaker := &Caretaker{}
	memento := caretaker.CreateMemento(*originator)
	fmt.Println(memento.GetState())
	if memento.GetState() != originator.state {
		t.Error("create memento error")
	}

	originator.SetState("change state")
	fmt.Println(originator.GetState())
	fmt.Println(memento.GetState())
	if memento.GetState() == originator.state {
		t.Error("change state error")
	}

	*originator = caretaker.RecoverOriginator(memento)
	fmt.Println(originator.GetState())
}
