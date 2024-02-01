package src

import (
	"errors"
	"fmt"
)

type Attendant struct {
	parkingLots []*ParkingLot
	strategy    Strategy
}

func NewAttendant(strategies ...Strategy) (*Attendant, error) {
	attendant := &Attendant{}
	attendant.parkingLots = []*ParkingLot{}
	if len(strategies) > 0 {
		attendant.strategy = strategies[0]
	} else {
		attendant.strategy = NEAREST
	}
	return attendant, nil
}

func (a *Attendant) assign(parkingLot *ParkingLot) {
	a.parkingLots = append(a.parkingLots, parkingLot)
}

func (a *Attendant) park(car *Car) (string, error) {
	parkingLotsToIterate := a.strategy.GetParkingLots(a.parkingLots)
	for i := range parkingLotsToIterate {
		ticket, err := parkingLotsToIterate[i].park(car, a.strategy)
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
