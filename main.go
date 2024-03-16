package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	bookings := []string{}
	var remainingTickets uint = 50
	const conferenceTickets int = 50

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for username
		fmt.Println("Please enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Please enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Please enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Please enter number of tickets you like to purchase: ")
		fmt.Scan(&userTickets)
		if userTickets != 0 {
			fmt.Printf("Thank you for purchasing %v tickets.\n", userTickets)
			bookings = append(bookings, firstName+" "+lastName)
		} else {
			fmt.Printf("Sorry, you did not purchase any ticket!")
		}
		if userTickets > remainingTickets {
			fmt.Printf("We have only %v tickets remaining. You can't book %v tickets. Please try again! \n", remainingTickets, userTickets)
			continue
		}
		remainingTickets = remainingTickets - userTickets

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
		fmt.Printf("We have total of %v tickets and %v tickets are still available.\n", conferenceTickets, remainingTickets)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("There first names of our bookings: %v\n", firstNames)

		if remainingTickets == 0 {
			// end program
			fmt.Printf("Our conference is booked out. Come back next year!")
			break
		} else {
			fmt.Printf("We have total of %v tickets and %v still available.\n", conferenceTickets, remainingTickets)
		}
	}
}
