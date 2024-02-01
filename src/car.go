package src

type Car struct {
	registrationNumber string
	color              CarColor
}

func NewCar(registrationNumber string, color CarColor) Car {
	return Car{registrationNumber: registrationNumber, color: color}
}
