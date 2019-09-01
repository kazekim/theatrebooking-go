/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package logic

import (
	"github.com/kazekim/theatrebooking-go/app/domain/entity"
	"github.com/kazekim/theatrebooking-go/app/domain/model"
	"math"
)

// SeatTicket find Ticket from seat index function
func SeatTicket(seat entity.Seat, numCol int, numRow int) model.Ticket {

	row := numRow - int(math.Floor(float64(seat.Index) / float64(numCol)))

	index := seat.Index % numCol
	index = seatNumberByIndex(index, numCol)

	ticket := model.Ticket{
		Row: row,
		SeatNumber: index,
		Name: seat.Name,
	}

	return ticket
}

// centerOf find center of row
func centerOf(numCol int) int {
	center := int(math.Ceil(float64(numCol)/2.0))

	if numCol %2 == 0 {
		center += 1
	}
	return center
}

// seatNumberByIndex find seat number from index
func seatNumberByIndex(index int, numCol int) int {

	centerIdx := centerOf(numCol)

	seatNumber := centerIdx

	idx := index - 1
	pos := int(math.Floor(float64(idx) / 2.0)) + 1


	if numCol %2 == 0 {
		if idx %2 == 0 {
			seatNumber -= pos
		}else{

			seatNumber += pos
		}
	}else{
		if idx %2 == 0 {
			seatNumber += pos
		}else{

			seatNumber -= pos
		}
	}
	return seatNumber
}