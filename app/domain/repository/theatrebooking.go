/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package repository

import "github.com/kazekim/theatrebooking-go/app/domain/entity"

type TheatreBookingRepository interface {
	DoBooking(name string) (*entity.Seat, error)
	GetNumRow() int
	GetNumSeatsPerRow() int
	SeatsLeft() int
}

