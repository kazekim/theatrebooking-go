/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package mock

import (
	"errors"
	"github.com/kazekim/theatrebooking-go/app/domain/entity"
)

type theatreBoookingRepository struct {
	col int
	row int
	seats []entity.Seat
}

func NewTheatreBookingRepository(col int, row int) *theatreBoookingRepository {
	var seats []entity.Seat
	return &theatreBoookingRepository{
		col: col,
		row: row,
		seats: seats,
	}
}

// DoBooking book a seat and store in database
func (r *theatreBoookingRepository) DoBooking(name string) (*entity.Seat, error) {

	if len(r.seats) == r.col*r.row {
		return nil, errors.New("The number of bookings are exceed the limit")
	}
	seat := entity.Seat {
		Name: name,
		Index: len(r.seats),
	}
	r.seats = append(r.seats, seat)

	return &seat, nil
}

// GetNumSeatsPerRow get number of seats per row
func (r *theatreBoookingRepository) GetNumSeatsPerRow() int {
	return r.col
}

// GetNumRow get number of row in theatre
func (r *theatreBoookingRepository) GetNumRow() int {
	return r.row
}

// SeatsLeft get number of seats left in theatre
func (r *theatreBoookingRepository) SeatsLeft() int {
	return r.col*r.row - len(r.seats)
}