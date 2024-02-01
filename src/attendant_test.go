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

func TestAssignOneParkingLotToMultipleAttendant(t *testing.T) {
	attendant, _ := NewAttendant()
	attendant1, _ := NewAttendant()
	parkingLot, _ := NewParkingLot(2)
	attendant.assign(parkingLot)
	attendant1.assign(parkingLot)
	car := NewCar("Ab", BLUE)
	car1 := NewCar("Ed", BLACK)

	attendant.park(car)
	attendant1.park(car1)

	if !parkingLot.isFull() {
		t.Errorf("Multiple attendant should be able to park in same parking lot.")
	}
}

func TestAttendantCanUnparkCarWhichIsParkByOtherAttendant(t *testing.T) {
	attendant, _ := NewAttendant()
	attendant1, _ := NewAttendant()
	parkingLot, _ := NewParkingLot(2)
	attendant.assign(parkingLot)
	attendant1.assign(parkingLot)
	car := NewCar("Ab", BLUE)

	ticket, _ := attendant.park(car)
	unparkedCar, _ := attendant1.unPark(ticket)

	if unparkedCar != car {
		t.Errorf("Attendant should be unpark car which parked by other attendant.")
	}
}

func TestParkCarNearestSlot(t *testing.T) {
	attendant, _ := NewAttendant(NEAREST)
	parkingLot, _ := NewParkingLot(1)
	parkingLot1, _ := NewParkingLot(1)
	car := NewCar("Ab", BLACK)
	attendant.assign(parkingLot)
	attendant.assign(parkingLot1)

	attendant.park(car)

	if !parkingLot.isFull() && parkingLot1.isFull() {
		t.Errorf("Car should be parked in parking lot.")
	}
}

func TestParkCarFarthestSlot(t *testing.T) {
	attendant, _ := NewAttendant(FARTHEST)
	parkingLot, _ := NewParkingLot(1)
	parkingLot1, _ := NewParkingLot(1)
	car := NewCar("Ab", BLACK)
	attendant.assign(parkingLot)
	attendant.assign(parkingLot1)

	attendant.park(car)

	if parkingLot.isFull() && !parkingLot1.isFull() {
		t.Errorf("Car should be parked in parking lot.")
	}
}

func TestParkCarsUsingDistributedStrategy(t *testing.T) {
	attendant, _ := NewAttendant(DISTRIBUTED)
	parkingLot, _ := NewParkingLot(2)
	parkingLot1, _ := NewParkingLot(2)
	car := NewCar("Ab", BLACK)
	car1 := NewCar("Ef", BLACK)
	attendant.assign(parkingLot)
	attendant.assign(parkingLot1)

	attendant.park(car)
	attendant.park(car1)

	if parkingLot.isFull() && parkingLot1.isFull() {
		t.Errorf("Cars should be park equaly in parking lots.")
	}
}
