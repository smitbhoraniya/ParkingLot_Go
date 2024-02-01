package src

import (
	"errors"
	"fmt"
)

type Attendant struct {
	parkingLots []*ParkingLot
}

func NewAttendant() (*Attendant, error) {
	attendant := &Attendant{}
	attendant.parkingLots = []*ParkingLot{}
	return attendant, nil
}

func (a *Attendant) assign(parkingLot *ParkingLot) {
	a.parkingLots = append(a.parkingLots, parkingLot)
}

func (a *Attendant) park(car *Car) (string, error) {
	for i := range a.parkingLots {
		ticket, err := a.parkingLots[i].park(car)
		if err != nil {
			fmt.Println(err.Error() + " trying other parking lots.")
		} else {
			return ticket, nil
		}
	}

	return "", errors.New("all parking lots are full")
}

func (a *Attendant) unPark(ticket string) (*Car, error) {
	for i := range a.parkingLots {
		car, err := a.parkingLots[i].unPark(ticket)
		if err != nil {
			fmt.Println(err.Error() + " trying other parking lots.")
		} else {
			return car, nil
		}
	}

	return nil, errors.New("car is not parked in any parking lot")
}
