package src

type Strategy int

const (
	NEAREST Strategy = iota
	FARTHEST
	DISTRIBUTED
)

type StrategyFunctions struct {
	getParkingLots func([]*ParkingLot) []*ParkingLot
	getSlot        func([]*Slot) *Slot
}

var strategies = map[Strategy]StrategyFunctions{
	NEAREST: {
		getParkingLots: func(parkingLots []*ParkingLot) []*ParkingLot {
			return parkingLots
		},
		getSlot: func(slots []*Slot) *Slot {
			for _, slot := range slots {
				if slot.isFree() {
					return slot
				}
			}
			return nil
		},
	},
	FARTHEST: {
		getParkingLots: func(parkingLots []*ParkingLot) []*ParkingLot {
			length := len(parkingLots)
			reversedLots := make([]*ParkingLot, length)

			for i, j := 0, length-1; i < length; i, j = i+1, j-1 {
				reversedLots[i] = parkingLots[j]
			}

			return reversedLots
		},
		getSlot: func(slots []*Slot) *Slot {
			for i := len(slots) - 1; i >= 0; i-- {
				if slots[i].isFree() {
					return slots[i]
				}
			}
			return nil
		},
	},
}

func (s Strategy) GetAvailableSlot(slots []*Slot) *Slot {
	return strategies[s].getSlot(slots)
}

func (s Strategy) GetParkingLots(parkingLots []*ParkingLot) []*ParkingLot {
	return strategies[s].getParkingLots(parkingLots)
}
