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

	ticket, err := slot.park(car)
	if err != nil && ticket != "" {
		t.Errorf(err.Error())
	}
}

func TestIsValidParkingTicket(t *testing.T) {
	car := NewCar("Ab", BLACK)
	slot := NewSlot()

	ticket, err := slot.park(car)
	if !slot.isValidTicket(ticket) && err == nil {
		t.Errorf("Ticket is not valid.")
	}
}

func TestShouldBeUnparkCar(t *testing.T) {
	car := NewCar("Ab", BLACK)
	slot := NewSlot()

	ticket, _ := slot.park(car)
	unparkedCar, err := slot.unPark(ticket)

	if car != unparkedCar && err == nil && !slot.isValidTicket(ticket) {
		t.Errorf("Invalid car unparked.")
	}
}

func TestUnparkCarWithEmptySlot(t *testing.T) {
	slot := NewSlot()

	car, err := slot.unPark("abc")

	if err == nil && car != nil {
		t.Errorf("Couldn't unpark empty slot.")
	}
}

func TestUnparkCarWithInvalidTicket(t *testing.T) {
	car := NewCar("Ab", BLACK)
	slot := NewSlot()

	slot.park(car)
	car, err := slot.unPark("abc")

	if err == nil && car != nil {
		t.Errorf("Couldn't unpark car using invalid ticket.")
	}
}
