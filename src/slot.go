package src

import (
	"errors"
	"os/exec"
)

type Slot struct {
	car    *Car
	ticket string
}

func NewSlot() *Slot {
	return &Slot{}
}

func (s *Slot) isFree() bool {
	return s.car == nil
}

func (s *Slot) park(car *Car) (string, error) {
	if !s.isFree() {
		return "", errors.New("Slot is occupied")
	}

	s.car = car
	newUUID, _ := exec.Command("uuidgen").Output()
	s.ticket = string(newUUID)
	return s.ticket, nil
}

func (s *Slot) isValidTicket(ticket string) bool {
	return s.ticket == ticket
}

func (s *Slot) unPark(ticket string) (*Car, error) {
	if s.isFree() {
		return nil, errors.New("Slot is empty")
	}
	if !s.isValidTicket(ticket) {
		return nil, errors.New("invalid ticket")
	}

	res := s.car
	s.car = nil
	s.ticket = ""
	return res, nil
}
