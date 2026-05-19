// always need to start wiht a package statement
package main

import "fmt"

// an entry point for the go program
func main() {
	var conferenceName = "Go Conference" // variable in Go
	const conferenceTickets int = 50  // constants in go
	var remainingTickets uint = 50

	fmt.Printf("conferenceTicket is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	// using printf
	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still availalbe\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")


	// ask user for their name
	var userName string
	var userTickets int
	
	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}
