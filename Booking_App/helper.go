package main

import (
	"strings"
)

func ValidateUserInput(firstname string, lastname string, email string, no_of_tickets uint, remaining_tickets uint) (bool, bool, bool) {

	var isValidName = len(firstname) >= 2 && len(lastname) >= 2

	isValidEmail := strings.Contains(email, "@")

	isValidTicket := no_of_tickets > 0 && no_of_tickets <= remaining_tickets
	return isValidName, isValidEmail, isValidTicket
}
