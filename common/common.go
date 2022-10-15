package common

import (
	"fmt"
	"strings"
)

func GreetUsers(confName string, remainingTickets uint, confTickets uint) {
	fmt.Printf("Welcome to %v, our booking application\n", confName)
	fmt.Printf("We have a total of %v out of %v\n", remainingTickets, confTickets)
	fmt.Println("Get your tickets here to attend")
}

func GetUserData(userName *string, userEmail *string) (bool, bool) {
	fmt.Println("Enter your userName: ")
	fmt.Scan(userName)

	fmt.Println("Enter your email: ")
	fmt.Scan(userEmail)

	isValidName := len(*userName) > 2
	isValidEmail := strings.Contains(*userEmail, "@")

	return isValidName, isValidEmail
}

func GetUserTickets(userTickets *uint, remainingTickets *uint) bool {
	fmt.Println("How many tickets? ")
	fmt.Scan(userTickets)

	return *userTickets > 0 && *userTickets <= *remainingTickets
}

func FinallizePurchase(userName string, userEmail string, userTickets uint, remainingTickets uint, confName string, bookings []string) {
	fmt.Printf("User %v with email %v booked %v tickets.\n", userName, userEmail, userTickets)

	fmt.Printf("%v tickets remaining for the %v\n", remainingTickets, confName)

	fmt.Printf("Current bookings are: %v\n", bookings)
}
