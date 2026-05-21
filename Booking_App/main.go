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

	// DATA STRUCTURES (ARRAYS AND SLICES IN GO)
	// var bookings = [50]string{"Nana", "Nicole", "Peter"}  // we defined the length/size of the array to be 50 (can have only 50 elements), we also specify the allowed data types to be strings

	// We can also declare an empty array with its type without assigning values yet
	// var bookings [50]string

	// SLICES IN GO
	/**
	Slice is an abstraction of an Array

	Slices are more flexible and powerful: variable-length or get a sub-array of its own

	Slices are also index-based and have a size, but is resized when needed
	*/

	// creating a slice
	var bookings []string

	// Alternative ways (to create a slice with values, use the go syntactic user for slice)
	// bookings2 := []string{"Nana", "Bob", "Segun"}
	// var bookings2 = []string{"Nana", "Bob", "Segun"}



	// adding elements to the array
	// bookings[0] = "Nana"
	// bookings[1] = "Nicole"

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
	fmt.Println("Please Enter your first name: ")
	fmt.Scan(&firstName) // scans user input and assigns to variable (uses the pointer operator & to store the address of the variable as received from the user)
	
	fmt.Println("Please Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Please Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Please Enter the number of tickets you wish to purchase: ")
	fmt.Scan(&userTickets)

	// Add this validation check before doing the subtraction!
	if userTickets > remainingTickets {
		fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets.\n", remainingTickets, userTickets)
		return // Stops the program so it doesn't execute the rest of the code
	}

	remainingTickets -= userTickets

	// add user data to the array
	// bookings[0] = firstName + " " + lastName

	// add data to SLICE (using the append function)
	bookings = append(bookings, firstName + " " + lastName)

	//ARRAYS
	// fmt.Printf("The whole array: %v\n", bookings)
	// fmt.Printf("The first value: %v\n", bookings[0])
	// fmt.Printf("Array type: %T\n", bookings[0])
	// fmt.Printf("Array length: %v\n", len(bookings))

	//SLICE
	// fmt.Printf("The whole slice: %v\n", bookings)
	// fmt.Printf("The first value: %v\n", bookings[0])
	// fmt.Printf("Slice type: %T\n", bookings)
	// fmt.Printf("Slice length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

	fmt.Printf("There are %v tickets remaining for %v\n", remainingTickets, conferenceName)

	fmt.Printf("These are all our bookings %v\n", bookings)

}
