package endpoint

import (
	"UmbrellaTask/pkg/tools"
	"encoding/json"
	"net/http"
)

// TODO more endpoints

type Service interface {
	DaysLeft() int
}

type Endpoint struct {
	service Service
}

func New(service Service) *Endpoint {
	return &Endpoint{service}
}

func (e *Endpoint) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	days := e.service.DaysLeft()

	resp, err := json.Marshal(map[string]int{
		"days": days,
	})

	if tools.MinorError(err, "Failed to marshal response") {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	_, err = w.Write(resp)
	tools.MinorError(err, "Failed to write response")
}
