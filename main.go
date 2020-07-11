package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	h "github.com/rinosukmandityo/message-flow/api"
)

func main() {
	r := h.RegisterHandler()
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./")))

	errs := make(chan error, 2)
	go func() {
		log.Printf("Listening on port %s\n", httpPort())
		errs <- http.ListenAndServe(httpPort(), r)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)

	}()
	log.Printf("Terminated %s", <-errs)

}

func httpPort() string {
	port := "24545"
	if os.Getenv("port") != "" {
		port = os.Getenv("port")
	}
	return fmt.Sprintf(":%s", port)
}
