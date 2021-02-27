package template

import (
	"testing"
)

func TestBoy_BeforeAction(t *testing.T) {
	boy := &Boy{}
	person := new(Person)

	person.SetName("boy")
	person.Concrete = boy
	boy.Person = *person

	person.Exit()
}
