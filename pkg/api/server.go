package api

import (
	"Car/pkg/app"
	"Car/pkg/repository"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Payload string `json:"payload"`
}

type Server struct {
	Router  *echo.Echo
	Methods app.ServerMethods
}

func New() *Server {
	router := echo.New()
	db := repository.Database{}
	serverMethods := app.ServerMethods{
		Storage: db,
	}

	server := Server{
		Router:  router,
		Methods: serverMethods,
	}
	return &server
}

func (s *Server) Run() error {
	return s.Router.Start(":8000")
}

func (s *Server) Shutdown(timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := s.Router.Shutdown(ctx); err != nil {
		//	log.Fatalf("shutdown timeout: %v", err)
		fmt.Printf("shutdown timeout: %v", err)
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

// Respond with give HTTP status code and value v.
// Always sends a JSON response, which means that v must be serializable as JSON.
func (s *Server) Respond(c echo.Context, code int, v interface{}) error {
	c.Response().Header().Set("User-Agent", "Christopher")
	return c.JSON(code, v)
}

// ErrBadRequest responds with a 400 error and given error message
func (s *Server) ErrBadRequest(message string) error {
	return echo.NewHTTPError(http.StatusBadRequest, message)
}
