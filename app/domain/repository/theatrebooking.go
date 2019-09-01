/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package repository

import "github.com/kazekim/theatrebooking-go/app/domain/entity"

// TheatreBookingRepository interface for TheatreBooking Repository
type TheatreBookingRepository interface {
	DoBooking(name string) (*entity.Seat, error)
	GetNumRow() int
	GetNumSeatsPerRow() int
	SeatsLeft() int
}

