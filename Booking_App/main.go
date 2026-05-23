// always need to start wiht a package statement
package main

import (
	"fmt"
	"strings"
)

// an entry point for the go program
func main() {
	conferenceName := "Go Conference" // syntactic sugar, := assignment in Go. PS: Does not work for variable with explicit type set, and does not work for constants
	const conferenceTickets int = 50  // constants in go
	var remainingTickets uint = 50

	// calling a function (all below have been replace by the function)
	greetUser(conferenceName, conferenceTickets, remainingTickets)

	// creating a slice
	var bookings []string
	var bookingsCount uint

	// LOOPS IN GO:
	for {
		// REPLACED WITH FUNCTIONS to get user inputs:
		firstName, lastName, email, userTickets := getUserInput()

		// REPLACED USER VALIDATIONS WITH FUNCTOINS
		isValidName, isValidEmail, isValidTicketNumber := validateUserInputs(firstName, lastName, email, userTickets, remainingTickets)

		// Add validation check before doing the subtraction, all have to be true before executing the booking
		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(remainingTickets, userTickets, bookings, int(bookingsCount), firstName, lastName, email, conferenceName)

			// replaced with functions to return firstNames
			firstNames := getFirstNames(bookings)

			// print firstNames
			fmt.Printf("The first names of bookings are: %v\n\n", firstNames)

			if remainingTickets == 0 {
				// END THE PROGRAM
				fmt.Printf("Go Conference is sold out! No more bookings available, come back next year!\n\n")
				fmt.Printf("Total Bookings: %v, names of invidividuals: %v\n", bookingsCount, firstNames)
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

		// HOW SWTICH STATEMENTS WORK IN Go
		//say we have a variable called city and we have cases for what should be executed for different kinds of city
		// city := "London"

		// switch city {
		// case "New york":
		// 	// execute code for booking New York conference tickets
		// case "Singpore":
		// 	//execute code for booking Singapore
		// case "Berlin":
		// 	// execute code for Berlin
		// case "Mexico City":
		// 	// execute code for Mexico

		// // we can also have cases for one or more cities to execute same conditions for
		// case "Hong Kong", "Turkey":
		// 	// Execute code for Hong Kong and Turkey
		// }
	}

}

// Functions in Go (also dealing with arugments in function)
func greetUser(confName string, confTickets int, remTickets uint) {
	fmt.Printf("Welcome to the %v booking application\n", confName)
	fmt.Printf("We have total of %v tickets and %v are still availalbe\n", confTickets, remTickets)
	fmt.Println("Get your tickets here to attend")
	fmt.Println("")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func validateUserInputs(firstName string, lastName string, email string, userTickets int, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && uint(userTickets) <= remainingTickets

	// Go allows us to return multiple values in a function
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, int) {
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

	return firstName, lastName, email, userTickets

}

func bookTicket(remainingTickets uint, userTickets int, bookings []string, bookingsCount int, firstName string, lastName string, email string, conferenceName string) {
	remainingTickets -= uint(userTickets)
	bookingsCount++

	// add data to SLICE (using the append function)
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

	fmt.Printf("There are %v tickets remaining for %v\n", remainingTickets, conferenceName)
}
