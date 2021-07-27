package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)


func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request) {
		log.Println("Hello world")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Error", http.StatusBadRequest)
		}
		fmt.Fprintf(rw, "Hello %v", string(d))
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("good bye")
	})
	
	http.ListenAndServe(":9090", nil)
}