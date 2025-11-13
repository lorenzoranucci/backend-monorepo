package main

import (
	"fmt"
	"net/http"

	"github.com/homeruntech/backend-monorepo/quote-domain/shared/go-lib/printer"
	http2 "github.com/homeruntech/backend-monorepo/shared/go-lib/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, printer.PrintMessage("Hello, World!!!"))
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/hello", helloWorldHandler)
	fmt.Println("Starting server on :8080...")
	http2.StartServer()
}
