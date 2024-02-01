package src

import "testing"

func TestValidCar(t *testing.T) {
	car := NewCar("Ab", BLACK)

	if car == nil {
		t.Errorf("Car should be created with parameters.")
	}
}
