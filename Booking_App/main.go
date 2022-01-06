package main

import (
	"fmt"
	"time"
)

var conf_name = "GO Conference"

const conf_tickets = 50

var remaining_tickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstname     string
	lastname      string
	email         string
	no_of_tickets uint
}

// var wg = sync.WaitGroup{} , Runnig App without for loop

func main() {

	greetUsers()

	for {

		firstname, lastname, email, no_of_tickets := GetUserIP()
		isValidName, isValidEmail, isValidTicket := ValidateUserInput(firstname, lastname, email, no_of_tickets, remaining_tickets)

		// isValidCity := city=="Singapore" || city == "London"

		if isValidTicket && isValidEmail && isValidName {

			BookTicket(no_of_tickets, firstname, lastname, email)

			// wg.Add(1)

			go SendTicket(no_of_tickets, firstname, lastname, email)

			var firstnames = getfnames()

			fmt.Printf("These all are our bookings: %v \n", firstnames)

			notickets := remaining_tickets == 0

			if notickets {
				// end program
				fmt.Println("ALL Tickets sold out !!")
				break

			}

		} else {
			if !isValidName {
				fmt.Println("First name or Last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered is not valid")
			}
			if !isValidTicket {
				fmt.Println("Number of tickets you entered is invalid")
			}

		}

	}
	// wg.Wait()

}

func greetUsers() {
	fmt.Println("WELCOME TO BOOKING.COM")
	fmt.Println("We are offering only Total of", conf_tickets, "tickets and", remaining_tickets, "are remaining")
	fmt.Printf("Get Your %v Tickets Here. \n", conf_name)
}

func getfnames() []string {
	firstnames := []string{}
	for _, booking := range bookings {

		firstnames = append(firstnames, booking.firstname)

	}
	return firstnames
}

func GetUserIP() (string, string, string, uint) {
	var firstname string
	var no_of_tickets uint
	var email string
	var lastname string

	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstname)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastname)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter no. of tickets: ")
	fmt.Scan(&no_of_tickets)

	return firstname, lastname, email, no_of_tickets
}

func BookTicket(no_of_tickets uint, firstname string, lastname string, email string) {
	remaining_tickets = remaining_tickets - no_of_tickets

	// create map for a user
	var UserData = UserData{
		firstname:     firstname,
		lastname:      lastname,
		email:         email,
		no_of_tickets: no_of_tickets,
	}

	bookings = append(bookings, UserData)

	fmt.Printf("List of Bookings is %v\n", bookings)

	fmt.Printf("Thank You %v %v for booking %v tickets, you will recieve the confirmation email at %v \n", firstname, lastname, no_of_tickets, email)
	fmt.Printf("Only %v ticktes are remaining now\n", remaining_tickets)
}

func SendTicket(no_of_tickets uint, firstname string, lastname string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v Tickets for %v\n%v", no_of_tickets, firstname, lastname)
	fmt.Println("################################")
	fmt.Printf("Sending Ticket:\n%v to email address %v\n", ticket, email)
	fmt.Println("################################")
	// wg.Done()
}
