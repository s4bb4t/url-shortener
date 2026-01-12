package main

import (
	"fmt"
	"os"

	"github.com/s4bb4t/url-shortener/internal/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		fatal(err)
	}

	fmt.Println(cfg)

	// TODO: init logger

	// TODO: init storage

	// TODO: init server

	// TODO: run server
}

func fatal(err error) {
	println(err.Error())
	os.Exit(1)
}
