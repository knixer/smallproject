package api

import (
	"Car/pkg/repository"
)

type testCar struct {
	getCar func(input string) (repository.CarStruct, error)
}

func (tc *testCar) GetCar(input string) (repository.CarStruct, error) {
	return tc.getCar(input)
}

type fixture struct {
	api     *Server
	Methods *testCar
}

func setTestFixture() *fixture {
	srv := New()
	test := &testCar{}
	srv.Methods.Storage = test
	srv.Routes()
	return &fixture{
		api:     srv,
		Methods: test,
	}
}
