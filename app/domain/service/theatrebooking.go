/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package service

import (
	"github.com/kazekim/theatrebooking-go/app/domain/entity"
	"github.com/kazekim/theatrebooking-go/app/domain/repository"
)

type TheatreBookingService struct {
	repo repository.TheatreBookingRepository
}

func NewTheatreBookingService(repo repository.TheatreBookingRepository) *TheatreBookingService {
	return &TheatreBookingService{
		repo: repo,
	}
}

// DoBooking book a seat service
func (s *TheatreBookingService) DoBooking(name string) (*entity.Seat, error) {
	return s.repo.DoBooking(name)
}
