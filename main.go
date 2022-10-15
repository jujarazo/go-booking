package main

import "fmt"

func main() {
	confName := "Go Conference"
	const confTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("ConfTickets is %T, remainingTickets is %T, conferenceNames is %T\n", confName, confTickets, remainingTickets)

	fmt.Printf("Welcome to %v, our booking application\n", confName)
	fmt.Printf("We have a total of %v out of %v\n", remainingTickets, confTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var userName string
		var userTickets uint

		fmt.Println("Enter your userName: ")
		fmt.Scan(&userName)

		fmt.Println("How many tickets? ")
		fmt.Scan(&userTickets)

		bookings = append(bookings, userName)

		remainingTickets = remainingTickets - userTickets

		fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

		fmt.Printf("%v tickets remaining for the %v\n", remainingTickets, confName)

		fmt.Printf("Current bookings are: %v\n", bookings)
	}

}
