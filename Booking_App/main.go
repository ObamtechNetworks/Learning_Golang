// always need to start wiht a package statement
package main

import "fmt"

// an entry point for the go program
func main() {
	var conferenceName = "Go Conference" // variable in Go
	const conferenceTickets = 50  // constants in go
	var remainingTickets = 50

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")


}
