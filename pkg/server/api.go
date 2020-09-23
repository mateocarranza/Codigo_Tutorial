package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type api struct {
	router http.Handler
}

//Server es la interface
type Server interface {
	Router() http.Handler
}

//New server crea un nuevo server
func New() Server {
	a := &api{}

	r := mux.NewRouter()
	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}
