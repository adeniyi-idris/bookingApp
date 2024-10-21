package main

import "strings"

func ValidateUserInput(userName string, lastName string, email string, userTicket uint, remainingTickets uint) (bool, bool, bool){
	isValid := len(userName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTicket > 0 && userTicket <= remainingTickets
	return isValid, isValidEmail, isValidTicketNumber
}

// to export a function capitalize the first letter of the name