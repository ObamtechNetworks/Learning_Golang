// always need to start wiht a package statement
package main

import "fmt"

// an entry point for the go program
func main() {
	conferenceName := "Go Conference" // syntactic sugar, := assignment in Go. PS: Does not work for variable with explicit type set, and does not work for constants
	const conferenceTickets int = 50  // constants in go
	var remainingTickets uint = 50

	fmt.Printf("conferenceTicket is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	// using printf
	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still availalbe\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	fmt.Printf("\n")


	var firstName string
	var lastName string
	var email string
	var userTickets int
	// ask user for their name
	fmt.Println("Please Enter your first name: ")
	fmt.Scan(&firstName) // scans user input and assigns to variable (uses the pointer operator & to store the address of the variable as received from the user)
	
	fmt.Println("Please Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Please Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Please Enter the number of tickets you wish to purchase: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
}
