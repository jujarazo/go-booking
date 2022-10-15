package main

import (
	"fmt"
	"jujarazo/booking-app/common"
	t "jujarazo/booking-app/types"
)

// Package level to make them Global they must start with capitalize letter
const confName = "Go Conference"
const confTickets = 50

var remainingTickets uint = 50
var bookings = []t.UserData{}

// type UserData struct {
// 	userName string
// 	email    string
// 	tickets  uint
// }

func main() {
	common.GreetUsers(confName, confTickets, remainingTickets)

	for {
		var userName string
		var userEmail string
		var userTickets uint = 10

		isValidName, isValidEmail := common.GetUserData(&userName, &userEmail)

		if !isValidName || !isValidEmail {
			fmt.Println("The name or email was invalid")
			continue
		}

		isValidAmount := common.GetUserTickets(&userTickets, &remainingTickets)

		if !isValidAmount {
			fmt.Printf("We only have %v tickets remaining, so you cant book %v tickets\n", remainingTickets, userTickets)
			continue
		}

		var userData = t.UserData{
			UserName: userName,
			Email:    userEmail,
			Tickets:  userTickets,
		}

		bookings = append(bookings, userData)

		remainingTickets = remainingTickets - userTickets

		// For loop example
		// for _, booking := range bookings {
		// 	// Splice example
		// 	// var names = string.Fields(booking)
		// 	fmt.Println(booking)
		// }

		common.FinallizePurchase(userName, userEmail, userTickets, remainingTickets, confName, bookings)

		isSoldOut := remainingTickets == 0
		if isSoldOut {
			fmt.Println("The tickets for the conference sold out!")
			break
		}
	}
}
