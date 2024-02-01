package src

type Slot struct {
	car Car
	id  string
}

func NewSlot() Slot {
	return Slot{}
}

func (s Slot) isFree() bool {
	var slot Slot
	return s == slot
}
