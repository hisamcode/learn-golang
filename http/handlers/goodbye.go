package handlers

import "net/http"

type GoodBye struct {}

func NewGoodBye() *GoodBye {
	return &GoodBye{}
}

func (g *GoodBye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("byee"))
}