package src

import "testing"

func TestValidParkingLot(t *testing.T) {
	parkingLot, err := NewParkingLot(2)

	if parkingLot == nil && err != nil {
		t.Errorf("Not able create the parking lot.")
	}
}

func TestInvalidParkingLot(t *testing.T) {
	parkingLot, err := NewParkingLot(-2)

	if err == nil && parkingLot != nil {
		t.Errorf("Not able create the parking lot.")
	}
}

func TestParkingLotShouldBeEmptyWhenCreated(t *testing.T) {
	parkingLot, _ := NewParkingLot(1)

	if parkingLot.isFull() {
		t.Errorf("Parking lot should be full.")
	}
}

func TestParkCarUsingParkingLot(t *testing.T) {
	parkingLot, _ := NewParkingLot(1)
	car := NewCar("Ab", WHITE)

	parkingLot.park(car)

	if !parkingLot.isFull() {
		t.Errorf("Parking lot should be full.")
	}
}

func TestParkCarInFullParkingLot(t *testing.T) {
	parkingLot, _ := NewParkingLot(1)
	car := NewCar("Ab", WHITE)
	car1 := NewCar("Ef", BLUE)

	parkingLot.park(car)
	_, err := parkingLot.park(car1)

	if err == nil {
		t.Errorf("Can not park car in full parking lot.")
	}
}

func TestUnParkCarUsingParkingLot(t *testing.T) {
	parkingLot, _ := NewParkingLot(1)
	car := NewCar("Ab", WHITE)

	ticket, _ := parkingLot.park(car)
	unparkedCar, _ := parkingLot.unPark(ticket)

	if car != unparkedCar {
		t.Errorf("Unparked car should be same as parked car.")
	}
}

func TestUnParkCarFromWrongParkingLot(t *testing.T) {
	parkingLot, _ := NewParkingLot(1)
	parkingLot1, _ := NewParkingLot(1)
	car := NewCar("Ab", WHITE)

	ticket, _ := parkingLot.park(car)
	_, err := parkingLot1.unPark(ticket)

	if err == nil {
		t.Errorf("Can not unpark car from wrong parking lot.")
	}
}

func TestUnParkCarUsingWrongTicket(t *testing.T) {
	parkingLot, _ := NewParkingLot(1)
	car := NewCar("Ab", WHITE)

	parkingLot.park(car)
	_, err := parkingLot.unPark("Abc")

	if err == nil {
		t.Errorf("Can not unpark car using wrong ticket.")
	}
}

func TestEmptySlotCount(t *testing.T) {
	parkingLot, _ := NewParkingLot(2)
	car := NewCar("Ab", WHITE)

	parkingLot.park(car)

	if parkingLot.emptySlotCount() != 1 {
		t.Errorf("empty slots count should be 1.")
	}
}
