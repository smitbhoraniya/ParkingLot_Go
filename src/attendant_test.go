package src

import "testing"

func TestValidAttendant(t *testing.T) {
	attendant, _ := NewAttendant()

	if attendant == nil {
		t.Errorf("Car should be created with parameters.")
	}
}

func TestParkCarUsingAttendant(t *testing.T) {
	attendant, _ := NewAttendant()
	parkingLot, _ := NewParkingLot(1)
	car := NewCar("Ab", BLACK)
	attendant.assign(parkingLot)

	attendant.park(car)

	if !parkingLot.isFull() {
		t.Errorf("Car should be parked in parking lot.")
	}
}

func TestAttendantShouldBeAbleParkInMultiple(t *testing.T) {
	attendant, _ := NewAttendant()
	parkingLot, _ := NewParkingLot(1)
	parkingLot1, _ := NewParkingLot(1)
	car := NewCar("Ab", BLACK)
	car1 := NewCar("Ef", BLACK)
	attendant.assign(parkingLot)
	attendant.assign(parkingLot1)

	attendant.park(car)
	attendant.park(car1)

	if !parkingLot.isFull() && !parkingLot1.isFull() {
		t.Errorf("Car should be parked in parking lot.")
	}
}

func TestCanNotParkCarIfAllParkingLotIsFull(t *testing.T) {
	attendant, _ := NewAttendant()
	parkingLot, _ := NewParkingLot(1)
	parkingLot1, _ := NewParkingLot(1)
	car := NewCar("Ab", BLACK)
	car1 := NewCar("Ef", BLACK)
	car2 := NewCar("Gh", BLACK)
	attendant.assign(parkingLot)
	attendant.assign(parkingLot1)

	attendant.park(car)
	attendant.park(car1)
	_, err := attendant.park(car2)

	if err == nil {
		t.Errorf("Can not park car if all parking lot is full")
	}
}

func TestUnparkCarUsingAttendant(t *testing.T) {
	attendant, _ := NewAttendant()
	parkingLot, _ := NewParkingLot(1)
	car := NewCar("Ab", BLACK)
	attendant.assign(parkingLot)

	ticket, _ := attendant.park(car)
	unparkedCar, _ := attendant.unPark(ticket)

	if car != unparkedCar {
		t.Errorf("Car should be unparked.")
	}
}

func TestCanNotUnparkCarUsingWrongTicket(t *testing.T) {
	attendant, _ := NewAttendant()
	parkingLot, _ := NewParkingLot(1)
	attendant.assign(parkingLot)

	_, err := attendant.unPark("Abc")

	if err == nil {
		t.Errorf("Can not un park car using wrong ticket.")
	}
}
