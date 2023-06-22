package endpoint

import (
	"UmbrellaTask/pkg/tools"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// TODO more endpoints

const shortForm = "2006-Jan-01"

type Service interface {
	DaysLeft() int
	ChangeDate(newDate time.Time) error
}

type Endpoint struct {
	service Service
}

func New(service Service) *Endpoint {
	return &Endpoint{service}
}

// example requests:
// curl -v -H "User-Role: admin"  "http://localhost:8080/home"
// curl -v  "http://localhost:8080/home"

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

// example request: curl -v -H "Content-Type: application/json" -H "User-Role: admin" -d '{"date": "2024-Feb-01"}' "http://localhost:8080/change_date"

func (e *Endpoint) ChangeDate(w http.ResponseWriter, r *http.Request) {

	type requestDate struct {
		Date string `json:"date"`
	}

	var result requestDate

	err := json.NewDecoder(r.Body).Decode(&result)
	if tools.MinorError(err, "Failed to decode request body") {
		fmt.Println(result)
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write(nil)
		tools.MinorError(err, "Failed to write response")
		return
	}

	parsedDate, err := time.Parse(shortForm, result.Date)
	if tools.MinorError(err, "Failed to parse date") {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write(nil)
		tools.MinorError(err, "Failed to write response")
		return
	}

	err = e.service.ChangeDate(parsedDate)
	if tools.MinorError(err, "Failed to change date") {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write(nil)
		tools.MinorError(err, "Failed to write response")
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(nil)
	tools.MinorError(err, "Failed to write response")
}
