/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package usecase

import (
	"errors"
	"fmt"
	"github.com/kazekim/theatrebooking-go/app/domain/logic"
	"github.com/kazekim/theatrebooking-go/app/domain/model"
	"github.com/kazekim/theatrebooking-go/app/domain/repository"
	"github.com/kazekim/theatrebooking-go/app/domain/service"
	"github.com/kazekim/theatrebooking-go/app/interface/mock"
)

type theatreBookingUsecase struct {
	service *service.TheatreBookingService
	repo repository.TheatreBookingRepository
}

func NewTheatreBookingUsecase(service *service.TheatreBookingService, repo repository.TheatreBookingRepository) *theatreBookingUsecase {
	return &theatreBookingUsecase{
		service: service,
		repo: repo,
	}
}

func BuildTheatreBookingUseCase(col int, row int) *theatreBookingUsecase {
	r := mock.NewTheatreBookingRepository(col, row)
	s := service.NewTheatreBookingService(r)
	u := NewTheatreBookingUsecase(s, r)

	return u
}

// DoBookingWithNumberofSeats do booking with number of seats
func (u *theatreBookingUsecase) DoBookingWithNumberofSeats(number_of_seat int) (*[]model.Ticket, error) {

	if number_of_seat > u.repo.SeatsLeft() {
		return nil, errors.New("Not enough Seats for your booking")
	}


	numCol := u.repo.GetNumSeatsPerRow()
	numRow := u.repo.GetNumRow()

	var tickets []model.Ticket
	var errs []error
	for i := 0; i < number_of_seat ; i++ {
		ticket, err := u.doBooking("Customer", numCol, numRow)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		tickets = append(tickets, *ticket)
	}

	if len(errs) > 0 {
		err := fmt.Errorf("There are %d booking success with %d errors occured", len(tickets), len(errs))
		return &tickets, err
	}

	return &tickets, nil
}

// doBooking do booking by name for each seat
func (u *theatreBookingUsecase) doBooking(name string, numCol, numRow int) (*model.Ticket, error) {

	seat, err := u.service.DoBooking(name)
	if err != nil {
		return nil, err
	}
	ticket := logic.SeatTicket(*seat, numCol, numRow)

	return &ticket, nil
}