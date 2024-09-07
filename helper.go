package main

import "strings"

func ValidateUserInput(firstName string, lastName string, emailName string, userTickets uint, remainingTickets uint) (bool, bool, bool){
	isValidName := len(firstName)>=2 && len(lastName)>=2
	isValidEmail := strings.Contains(emailName, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}