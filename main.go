package main

import (
	"fmt"
	"strings"
)

func main() {
	confName := "Go Conference"
	const confTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}

	greetUsers(confName, confTickets, remainingTickets)

	for {
		var userName string = "holis"
		var userEmail string
		var userTickets uint = 10

		isValidName, isValidEmail := getUserData(&userName, &userEmail)

		if !isValidName || !isValidEmail {
			fmt.Println("The name or email was invalid")
			continue
		}

		isValidAmount := getUserTickets(&userTickets, &remainingTickets)

		if !isValidAmount {
			fmt.Printf("We only have %v tickets remaining, so you cant book %v tickets\n", remainingTickets, userTickets)
			continue
		}

		bookings = append(bookings, userName)

		remainingTickets = remainingTickets - userTickets

		// For loop example
		// for _, booking := range bookings {
		// 	// Splice example
		// 	// var names = string.Fields(booking)
		// 	fmt.Println(booking)
		// }

		fmt.Printf("User %v with email %v booked %v tickets.\n", userName, userEmail, userTickets)

		fmt.Printf("%v tickets remaining for the %v\n", remainingTickets, confName)

		fmt.Printf("Current bookings are: %v\n", bookings)

		isSoldOut := remainingTickets == 0
		if isSoldOut {
			fmt.Println("The tickets for the conference sold out!")
			break
		}
	}
}

func greetUsers(confName string, remainingTickets uint, confTickets uint) {
	fmt.Printf("Welcome to %v, our booking application\n", confName)
	fmt.Printf("We have a total of %v out of %v\n", remainingTickets, confTickets)
	fmt.Println("Get your tickets here to attend")
}

func getUserData(userName *string, userEmail *string) (bool, bool) {
	fmt.Println("Enter your userName: ")
	fmt.Scan(userName)

	fmt.Println("Enter your email: ")
	fmt.Scan(userEmail)

	isValidName := len(*userName) > 2
	isValidEmail := strings.Contains(*userEmail, "@")

	return isValidName, isValidEmail
}

func getUserTickets(userTickets *uint, remainingTickets *uint) bool {
	fmt.Println("How many tickets? ")
	fmt.Scan(userTickets)

	return *userTickets > 0 && *userTickets <= *remainingTickets
}
