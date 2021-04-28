package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) handleCarGet() echo.HandlerFunc {
	return func(c echo.Context) error {

		brand := c.Param("brand")

		car, err := s.Methods.GetCar(brand)

		if err != nil {
			return s.ErrBadRequest(fmt.Sprintf("could not find car brand: %v", err))
		}

		return s.Respond(c, http.StatusOK, car)
	}
}
