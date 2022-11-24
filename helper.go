package main

import "strings"

func validate(user_email string, user_tickets uint, remaining_tickets uint) (bool, bool) {
	isValidEmail := strings.Contains(user_email, "@")
	isValidTicket := user_tickets <= remaining_tickets && remaining_tickets > 0
	return isValidEmail, isValidTicket
}
