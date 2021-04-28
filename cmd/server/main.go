package main

import (
	"Car/pkg/api"
	"fmt"
	"net/http"

	"os"
	"os/signal"
	"syscall"
	"time"
)

func index_handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Whoa, Go is neat!!!")
}

func main() {
	//Ctrl+Alt+M to stop server

	server := api.New()

	server.Routes()

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := server.Run(); err != nil {
			fmt.Printf("listen: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	server.Shutdown(5 * time.Second)

}
