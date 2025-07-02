package service

import "time"

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) DaysLeft() int16 {
	d := time.Date(2025, time.November, 10, 0, 0, 0, 0, time.UTC)
	dur := time.Until(d)
	return int16(dur.Hours() / 24)
}
