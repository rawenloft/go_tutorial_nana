package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	var remainingTickets uint = 50
	const conferenceTickets int = 50

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	var bookings [50]string
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
		bookings[0] = firstName + " " + lastName
	} else {
		fmt.Printf("Sorry, you did not purchase any ticket!")
	}

	remainingTickets = remainingTickets - userTickets
	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Array type: %T\n", bookings)
	fmt.Printf("Array length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("We have total of %v tickets and %v tickets are still available.\n", conferenceTickets, remainingTickets)

}
