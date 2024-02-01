package src

import "testing"

func TestValidSlot(t *testing.T) {
	slot := NewSlot()

	if slot == nil {
		t.Errorf("Car should be created with parameters.")
	}
}

func TestSlotShouldBeEmptyWhenCreated(t *testing.T) {
	slot := NewSlot()

	if !slot.isFree() {
		t.Errorf("Slot should be free, when created.")
	}
}

func TestShouldBeParkCarInEmptySlot(t *testing.T) {
	car := NewCar("Ab", BLACK)
	slot := NewSlot()

	_, err := slot.park(car)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestIsValidParkingTicket(t *testing.T) {
	car := NewCar("Ab", BLACK)
	slot := NewSlot()

	ticket, _ := slot.park(car)
	if !slot.isValidTicket(ticket) {
		t.Errorf("Ticket is not valid.")
	}
}

func TestShouldBeUnparkCar(t *testing.T) {
	car := NewCar("Ab", BLACK)
	slot := NewSlot()

	ticket, _ := slot.park(car)
	unparkedCar, _ := slot.unPark(ticket)

	if car != unparkedCar {
		t.Errorf("Invalid car unparked.")
	}
}

func TestUnparkCarWithEmptySlot(t *testing.T) {
	slot := NewSlot()

	_, err := slot.unPark("abc")

	if err == nil {
		t.Errorf("Couldn't unpark empty slot.")
	}
}

func TestUnparkCarWithInvalidTicket(t *testing.T) {
	car := NewCar("Ab", BLACK)
	slot := NewSlot()

	slot.park(car)
	_, err := slot.unPark("abc")

	if err == nil {
		t.Errorf("Couldn't unpark car using invalid ticket.")
	}
}
