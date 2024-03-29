/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package logic

import (
	"github.com/kazekim/theatrebooking-go/app/domain/entity"
	"testing"
)

// TestSeatTicket Test case for find Ticket from seat index
func TestSeatTicket(t *testing.T) {

	//Test Set 1

	doTestSeatTicketWithData(t, 10, 6, 2, 10, 5)


	//Test Set 2

	doTestSeatTicketWithData(t, 10, 5, 19, 7, 1)


	//Test Set 3

	doTestSeatTicketWithData(t, 10, 20, 45, 8, 8)

}

// doTestSeatTicketWithData  find Ticket from seat index test function
func doTestSeatTicketWithData(t *testing.T, numRow, numCol, index, expectedRow, expectedCol int) {
	seat := entity.Seat{
		Name: "Test",
		Index: index,
	}

	ticket := SeatTicket(seat, numCol, numRow)

	if ticket.Row != expectedRow || ticket.SeatNumber != expectedCol {
		t.Errorf("Error! Test Seat for Index %d, Expected the result of %d, %d but instead got %d, %d", seat.Index, expectedRow, expectedCol, ticket.Row, ticket.SeatNumber)
	}
}

