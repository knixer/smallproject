package api

import (
	"Car/cmd/servertest"
	"Car/pkg/repository"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/matryer/is"
)

func TestGetThingsAll(t *testing.T) {
	is := is.NewRelaxed(t)
	fx := setTestFixture()

	car := repository.CarStruct{
		Model:           "SAAB 95",
		ModelYear:       2000,
		Gearbox:         "Manual",
		OdometerReading: 105000,
		Color:           "White",
	}

	fx.Methods.getCar = func(input string) (repository.CarStruct, error) {
		return car, nil
	}

	r := servertest.Get(fx.api, "/car/Saab")
	is.Equal(r.Code, http.StatusOK)

	var resultCar repository.CarStruct
	err := json.NewDecoder(r.Body).Decode(&resultCar)

	is.NoErr(err)
	is.Equal(car, resultCar)

}
