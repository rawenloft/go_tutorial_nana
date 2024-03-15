package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	var remainingTickets = 50
	const conferenceTickets = 50

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	var userName string
	var userTickets int
	// ask user for username
	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

}
