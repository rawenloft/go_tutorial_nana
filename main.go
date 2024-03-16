package main

import (
	"fmt"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName     string
	lastName      string
	email         string
	bookedTickets uint
}

func main() {

	greetUsers()

	for {
		//getUserInput
		firstName, lastName, email, userTickets := getUserInput()

		// validateInput
		isValidName, isValidEmail, isValidTicketsNumber := ValidateInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketsNumber {
			fmt.Printf("Thank you for purchasing %v tickets.\n", userTickets)

		} else {
			fmt.Printf("Sorry, you did not purchase any ticket!\n")
		}

		if isValidName && isValidEmail && isValidTicketsNumber {
			// bookTickets
			bookTicket(userTickets, firstName, lastName, email)
			// sendTicket
			go sendTicket(userTickets, firstName, lastName, email)
			fmt.Printf("There first names of our bookings: %v\n", getFirstNames())

			if remainingTickets == 0 {
				// end program
				fmt.Printf("Our conference is booked out. Come back next year!")
				break
			}
		} else {

			if !isValidName {
				fmt.Println("You enter invalid name. First name or last name is too short.")
			}
			if !isValidEmail {
				fmt.Println("You entered invalid email.")
			}
			if !isValidTicketsNumber {
				fmt.Println("You entered invalid number of tickets.")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for information
	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Please enter number of tickets you like to purchase: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create map for user
	var userData = UserData{
		firstName:     firstName,
		lastName:      lastName,
		email:         email,
		bookedTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are still available for %v.\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(3 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket %v to email address %v\n", ticket, email)
	fmt.Println("#################")

}
