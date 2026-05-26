// always need to start with a package statement
package main

import (
	"booking_app/helper"
	"fmt"
)

// PACKAGE-LEVEL VARIABLES
// These are defined outside main, so ALL functions below can share and modify them directly!
// NOTE: WE CANNOT CREATE PACKAGE LEVEL VARIABLES USING THE := syntax, we must explicitly use the var keyword
var conferenceName = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0) // converted to a struct type
var bookingsCount uint

// STRUCT DATA TYPE IN GO
// in Go, struct is like creating a custom type that can hold multiple fields of different data types, it's like a blueprint for creating objects that can represent complex data structures. We can define a struct to represent a user booking, for example, with fields for first name, last name, email, number of tickets, and whether they opted in for the newsletter. This allows us to organize and manage our data more effectively.
type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets int
}

func main() {

	// the greet user function
	greetUser()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInputs(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			// We don't need to pass remainingTickets or bookings anymore, they can be accessed from the package level
			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()

			fmt.Printf("The first names of bookings are: %v\n\n", firstNames)

			// This will now correctly evaluate to true when tickets hit 0
			if remainingTickets == 0 {
				fmt.Printf("Go Conference is sold out! No more bookings available, come back next year!\n\n")
				fmt.Printf("Total Bookings: %v, names of individuals: %v\n", bookingsCount, firstNames)
				break
			}
		} else {
			if !isValidName {
				fmt.Printf("\nFirst name or Last Name is too short\n")
			}
			if !isValidEmail {
				fmt.Printf("The Email address you entered is invalid '@' sign missing\n")
			}
			if !isValidTicketNumber {
				fmt.Printf("The Number of tickets is invalid. We only have %v tickets remaining\n\n", remainingTickets)
			}
		}
	}
}

// Notice we stripped out parameters that are now global variables!
func greetUser() {
	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	fmt.Println("")
}

func getFirstNames() []string {
	firstNames := []string{}

	// now that we are using map, accessing first names will be way more easier
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Please Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Please Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Please Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Please Enter the number of tickets you wish to purchase: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets int, firstName string, lastName string, email string) {
	// Modifying these here alters the shared global variables permanently
	remainingTickets -= uint(userTickets)
	bookingsCount++

	// Now dealing with structs => Create a struct type for the user data and assign values to the function arguments
	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings) // prints the lists of maps, basically an array of maps (objects)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("There are %v tickets remaining for %v\n", remainingTickets, conferenceName)
}
