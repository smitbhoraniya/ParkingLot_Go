package src

import "testing"

func TestCarValid(t *testing.T) {
	car := NewCar("Ab", BLACK)
	car1 := Car{}

	if car == car1 {
		t.Errorf("Car should be created with parameters.")
	}
}
