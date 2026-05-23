// Define the package this file belongs to, in this case main
package helper

import "strings"

func ValidateUserInputs(firstName string, lastName string, email string, userTickets int, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && uint(userTickets) <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
