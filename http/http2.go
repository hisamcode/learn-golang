package main

import (
	"net/http"
)


func main() {
	// l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// gh := handlers.NewGoodbye(l)

	// sm := http.NewServeMux()
	// sm.Handle("/goodbye", gh)

	http.ListenAndServe(":9090", nil)
}