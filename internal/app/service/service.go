package service

import (
	"UmbrellaTask/pkg/tools"
	"errors"
	"time"
)

type Service struct {
	// TODO read from file/config
	bigDay time.Time
}

func New() *Service {
	location, err := time.LoadLocation("Asia/Novosibirsk")
	if tools.MinorError(err, "Failed to load location") {
		location = time.UTC
	}
	defaultDay := time.Date(2025, time.January, 1, 0, 0, 0, 0, location)
	return &Service{defaultDay}
}

func (s *Service) DaysLeft() int {

	duration := time.Until(s.bigDay)
	return int(duration.Hours() / 24)
}

func (s *Service) ChangeDate(newDate time.Time) error {
	todayDate := time.Now()
	if newDate.After(todayDate) {
		s.bigDay = newDate
		return nil
	}
	return errors.New("new date has already passed")
}
