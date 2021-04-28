package repository

import "errors"

type Database struct {
}

type CarStruct struct {
	Model           string
	ModelYear       int
	Gearbox         string
	OdometerReading int
	Color           string
}

var cars = map[string]CarStruct{
	"Tesla": CarStruct{
		Model:           "Model S Long Range AWD",
		ModelYear:       2020,
		Gearbox:         "Automatic",
		OdometerReading: 20250,
		Color:           "White",
	},
	"Volvo": CarStruct{
		Model:           "Volvo V70",
		ModelYear:       2007,
		Gearbox:         "Manual",
		OdometerReading: 212250,
		Color:           "Light Grey",
	},
	"Audi": CarStruct{
		Model:           "Audi A4",
		ModelYear:       2011,
		Gearbox:         "Manual",
		OdometerReading: 126990,
		Color:           "Black",
	},
	"BMW": CarStruct{
		Model:           "BMW 3-serien",
		ModelYear:       2007,
		Gearbox:         "Automatic",
		OdometerReading: 218910,
		Color:           "Dark Grey",
	},
	"Citroen ": CarStruct{
		Model:           "Citroen C3",
		ModelYear:       2015,
		Gearbox:         "Manual",
		OdometerReading: 30130,
		Color:           "White",
	},
}

func (db Database) GetCar(query string) (CarStruct, error) {

	car, ok := cars[query]
	if !ok {
		return CarStruct{}, errors.New("Car does not exist")
	}

	return car, nil

}
