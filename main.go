package main

import "fmt"

func main() {
	conferenceName := "Go Study Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	// varing Prints to check how do they work
	fmt.Printf("Welcome to %v booking system!\n", conferenceName)
	fmt.Println("Out of", conferenceTickets, "total tickets, we have", remainingTickets, "remaining")
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address:")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	remainingTickets -= uint(userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaingin for %v\n", remainingTickets, conferenceName)
}
