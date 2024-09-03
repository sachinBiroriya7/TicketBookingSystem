package main

import "strings"

func validateUserInput(tickets_needed, Total_tickets, remaining_tickets uint, FName, LName, email string) (bool, bool, bool, bool) {
	isValidFName := len(FName) >= 2
	isValidLName := len(LName) >= 2
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isvalidTickets := tickets_needed <= Total_tickets && tickets_needed <= remaining_tickets

	return isValidEmail, isValidFName, isValidLName, isvalidTickets
}
