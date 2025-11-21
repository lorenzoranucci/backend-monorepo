package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/caarlos0/env/v11"
	"github.com/lorenzoranucci/backend-monorepo/quote-domain/shared/go-lib/printer"
	http2 "github.com/lorenzoranucci/backend-monorepo/shared/go-lib/http"
)

type Config struct {
	Version string `env:"Version"`
}

func main() {
	cfg := &Config{}
	if err := env.ParseWithOptions(cfg, env.Options{}); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/hello", helloWorldHandler)
	fmt.Println("Starting server on :8080...")
	http2.StartServer()
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, printer.PrintMessage("Hello, World!!!"))
	if err != nil {
		return
	}
}
