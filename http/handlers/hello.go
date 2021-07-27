package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Hello struct {}

func (h Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// read r.body
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Terjadi ke eroran!", http.StatusBadRequest)
	}

	// ubah r.body ke string
	d := string(b)

	// pesan dari dalem
	fmt.Printf("path / telah di eksekusi dengan data %v", d)

	// pesan ke luar
	fmt.Fprintf(rw, "hello %v", d)
}