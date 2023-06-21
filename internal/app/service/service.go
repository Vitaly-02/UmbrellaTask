package service

import (
	"UmbrellaTask/pkg/tools"
	"time"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) DaysLeft() int {
	location, err := time.LoadLocation("Asia/Novosibirsk")
	if tools.MinorError(err, "Failed to load location") {
		location = time.UTC
	}

	bigDay := time.Date(2025, time.January, 1, 0, 0, 0, 0, location)
	duration := time.Until(bigDay)
	return int(duration.Hours() / 24)
}
