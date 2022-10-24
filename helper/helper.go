package helper

import "strings"

func ValidateUserInput(FirstName string, LastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(FirstName) >= 2 && len(LastName) >= 2

	isValidEmail := strings.Contains(email, "@")

	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber

}

// func ValidateUserInput()(bool, bool, bool){
// 	isValidName := len(FirstName) >= 2 && len(LastUsername) >= 2

// 	isValidPassword := strings.Contains(email, @)

// }
