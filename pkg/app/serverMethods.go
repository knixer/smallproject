package app

import (
	"Car/pkg/repository"
)

type Storage interface {
	GetCar(input string) (repository.CarStruct, error)
}

type ServerMethods struct {
	Storage Storage
}

func (sm *ServerMethods) GetCar(input string) (repository.CarStruct, error) {
	return sm.Storage.GetCar(input)
}
