package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

var conference_name = "GO Conference"

const conference_tickets = 50

var remaining_tickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	name    string
	email   string
	tickets uint
}

var wg = sync.WaitGroup{}

func main() {

	// fmt.Printf("conference ticket is %T, remaining ticket is %T, conference is %T\n", conference_tickets, remaining_tickets, conference_name)
	greetUsers(conference_name, conference_tickets, remaining_tickets)

	for remaining_tickets > 0 {
		user_name, user_email, user_tickets := getUserInputs()
		if user_tickets > remaining_tickets {
			fmt.Printf("Only %v tickets remain. You can't get %v", remaining_tickets, user_tickets)
			continue
		}

		isValidEmail, isValidTicket := validate(user_email, user_tickets, remaining_tickets)

		if !isValidEmail || !isValidTicket {
			if !isValidEmail {
				fmt.Printf("Invalid Email Address")
			}
			if !isValidTicket {
				fmt.Printf("All the tickets are gone! Come back later!")
			}
			break
		}
		bookTickets(user_tickets, user_name, user_email)
		wg.Add(1)
		go sendTicket(user_tickets, user_name, user_email)
		firstNames := printNames()
		fmt.Printf("These are the booking list %v\n:", firstNames)
		wg.Wait()
	}

}
func greetUsers(conference_name string, conference_tickets uint, remaining_tickets uint) {
	fmt.Printf("Welcome to %v booking application!\n", conference_name)
	fmt.Printf("We have total of %v tickets and we have %v tickets still available.\n", conference_tickets, remaining_tickets)
	fmt.Println("Get your TICKETS here now!")
}

func printNames() []string {
	var firstNames []string
	for _, booking := range bookings {
		var firstName string = strings.Split(booking.name, " ")[0]
		firstNames = append(firstNames, firstName)
	}
	return firstNames
}

func getUserInputs() (string, string, uint) {
	scanner := bufio.NewScanner(os.Stdin)
	var user_name string
	var user_email string
	var user_tickets uint
	//ask user name
	fmt.Println("What's your name?")
	// fmt.Scan(&user_name)
	if scanner.Scan() {
		user_name = scanner.Text()
	}

	fmt.Println("What's your email?")
	fmt.Scan(&user_email)

	fmt.Println("How many tickets would you like?")

	fmt.Scan(&user_tickets)

	return user_name, user_email, user_tickets
}

func bookTickets(user_tickets uint, user_name string, user_email string) {
	remaining_tickets -= user_tickets
	var userData = UserData{
		name:    user_name,
		email:   user_email,
		tickets: user_tickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v for booking %v tickets.\nConfirmation mail will be sent to email address : %v\n", user_name, user_tickets, user_email)
	fmt.Println(remaining_tickets, "tickets remaining for the\n", conference_name)
}

func sendTicket(user_tickets uint, user_name string, user_email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v", user_tickets, user_name)
	fmt.Println("##############")
	fmt.Printf("Sending %v ticket:\n to email address %v\n", ticket, user_email)
	fmt.Println("##############")
	wg.Done()
}
