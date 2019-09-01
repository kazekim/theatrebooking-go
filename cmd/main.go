/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package main

import (
	"fmt"
	"github.com/kazekim/theatrebooking-go/app/usecase"
	"github.com/kazekim/theatrebooking-go/pkg/jhfmt"
)

func main() {

	col := 20
	row := 10

	u := usecase.BuildTheatreBookingUseCase(col, row)

	tickets, err := u.DoBookingWithNumberofSeats(1)
	if err != nil {
		fmt.Println(err)
	}

	jhfmt.PrintBeautyStruct(tickets)

	tickets, err = u.DoBookingWithNumberofSeats(3)
	if err != nil {
		fmt.Println(err)
	}

	jhfmt.PrintBeautyStruct(tickets)

	tickets, err = u.DoBookingWithNumberofSeats(6)
	if err != nil {
		fmt.Println(err)
	}

	jhfmt.PrintBeautyStruct(tickets)


	tickets, err = u.DoBookingWithNumberofSeats(20)
	if err != nil {
		fmt.Println(err)
	}

	jhfmt.PrintBeautyStruct(tickets)

}
