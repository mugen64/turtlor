package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/mugen64/turtlor/api"
	"github.com/mugen64/turtlor/configs"
	"github.com/mugen64/turtlor/pkg/logger"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	config, err := configs.LoadConfig()
	log := logger.NewLogger(config.LogLevel)
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
		return
	}
	errx := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errx <- fmt.Errorf("%s", <-c)
	}()
	go func() {
		log.Print("starting server on -> " + config.Server.AddressWithProtocol())
		handler := api.NewServer(config, log)
		errx <- http.ListenAndServe(config.Server.Address(), handler)
	}()
	log.PrintError(<-errx)

}
