package server

import (
	"UmbrellaTask/internal/app/endpoint"
	"UmbrellaTask/internal/app/mw"
	"UmbrellaTask/internal/app/service"
	"log"
	"net/http"
)

const address = "localhost:8080"

type Server struct {
	e    *endpoint.Endpoint
	serv *service.Service
	mux  *http.ServeMux
	// TODO config
	// TODO limiter
}

func New() (*Server, error) {
	s := &Server{}

	s.serv = service.New()
	s.e = endpoint.New(s.serv)

	var mux *http.ServeMux
	s.mux = mux

	return s, nil
}

func (s *Server) initHandlers() *http.ServeMux {
	mux := http.NewServeMux()

	homeHandler := http.HandlerFunc(s.e.Home)
	changeHandler := http.HandlerFunc(s.e.ChangeDate)

	mux.Handle("/change_date", mw.AdminHandler(changeHandler))
	mux.Handle("/home", mw.RoleCheck(homeHandler))

	return mux
}

func (s *Server) Start() {
	s.mux = s.initHandlers()
	log.Fatal(http.ListenAndServe(address, s.mux))
}
