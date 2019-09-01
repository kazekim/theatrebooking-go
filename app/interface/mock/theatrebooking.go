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

func (r *theatreBoookingRepository) GetNumSeatsPerRow() int {
	return r.col
}

func (r *theatreBoookingRepository) GetNumRow() int {
	return r.row
}

func (r *theatreBoookingRepository) SeatsLeft() int {
	return r.col*r.row - len(r.seats)
}