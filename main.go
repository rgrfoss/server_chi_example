package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rgrfoss/webserver/pkg/db"
	"github.com/rgrfoss/webserver/pkg/http/rest"
)

var (
	addr = ":8888"
)

func main() {

	// gracefully exit on keyboard interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	conn, err := db.PostgresConnect()
	if err != nil {
		os.Exit(1)
	}

	router := rest.InitHandlers(&rest.HandlerArgs{
		Conn: conn,
	})

	go func() {
		if err := http.ListenAndServe(addr, router); err != nil {
			os.Exit(1)
		}
	}()

	<-c
	os.Exit(0)
}
