package service

import "github.com/jafsq5/kea-tool-ui/internal/model"

type ReservationService struct{}

func NewReservationService() *ReservationService {
	return &ReservationService{}
}

func (s *ReservationService) List() []model.Reservation {
	return []model.Reservation{
		{
			MAC:      "52:54:00:11:22:33",
			Hostname: "server01",
			IP:       "192.168.10.10",
		},
		{
			MAC:      "52:54:00:AA:BB:CC",
			Hostname: "printer",
			IP:       "192.168.10.20",
		},
	}
}
