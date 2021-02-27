package builder

import (
	"fmt"
	"testing"
)

func TestBuilderPattern(t *testing.T) {
	director := Director{}

	car := &Car{}
	director.Construct(car)
	vehicle := director.builder.GetVehicle()

	fmt.Println(vehicle)
	if vehicle.Wheels != 4 {
		t.Errorf("car wheels must be 4, but get %d\n", vehicle.Wheels)
	}
	if vehicle.Seats != 4 {
		t.Errorf("car seats must be 4, but get %d\n", vehicle.Seats)
	}
}
