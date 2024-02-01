package src

import "testing"

func TestIsSlotFree(t *testing.T) {
	slot := NewSlot()

	if !slot.isFree() {
		t.Errorf("Slot should be free, when created.")
	}
}
