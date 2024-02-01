package src

import "errors"

type ParkingLot struct {
	slots []*Slot
}

func NewParkingLot(numberOfSlots int) (*ParkingLot, error) {
	if numberOfSlots < 0 {
		return nil, errors.New("number of slots should be positive")
	}

	parkingLot := &ParkingLot{}
	parkingLot.slots = make([]*Slot, numberOfSlots)
	for i := range parkingLot.slots {
		parkingLot.slots[i] = NewSlot()
	}
	return parkingLot, nil
}

func (p *ParkingLot) park(car *Car, strategies ...Strategy) (string, error) {
	var strategy Strategy
	if len(strategies) > 0 {
		strategy = strategies[0]
	} else {
		strategy = NEAREST
	}

	slot := strategy.GetAvailableSlot(p.slots)
	if slot == nil {
		return "", errors.New("parking lot is full")
	}

	ticket, err := slot.park(car)
	if err != nil {
		return "", errors.New(err.Error())
	}
	return ticket, nil
}

func (p *ParkingLot) isFull() bool {
	for i := range p.slots {
		if p.slots[i].isFree() {
			return false
		}
	}
	return true
}

func (p *ParkingLot) unPark(ticket string) (*Car, error) {
	for i := range p.slots {
		if p.slots[i].isValidTicket(ticket) {
			car, err := p.slots[i].unPark(ticket)
			if err != nil {
				return nil, errors.New(err.Error())
			}
			return car, nil
		}
	}

	return nil, errors.New("car is not parked in this parking lot")
}

func (p *ParkingLot) emptySlotCount() int {
	count := 0
	for i := range p.slots {
		if p.slots[i].isFree() {
			count++
		}
	}

	return count
}
