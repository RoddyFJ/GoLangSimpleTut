package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName    string
	lastName     string
	emailAddress string
	numTickets   uint
}

var wg = sync.WaitGroup{}

func main() {

	greatUsers()

	for {

		firstName, lastName, emailAddress, userTickets := getUserInfo()
		isValidName, isValidEmail, isValidTickets := ValidateUserInput(firstName, lastName, emailAddress, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidTickets {

			bookTicket(userTickets, firstName, lastName, emailAddress)
			wg.Add(1)
			go sendTickets(userTickets, firstName, lastName, emailAddress)

			firstNames := getFirstNames()

			fmt.Printf("These are all the names for bookings: %v\n", firstNames)
			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				//end program
				fmt.Println("Our conference is sold out")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Your first or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Your email address must have a @")
			}
			if !isValidTickets {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}
		wg.Wait()
	}

}

func greatUsers() {
	fmt.Printf("Welcome to our %v conference\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v left\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your Tickets ->\n")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {

		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInfo() (string, string, string, uint) {
	var firstName string
	var lastName string
	var emailAddress string
	var userTickets uint
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address: ")
	fmt.Scan(&emailAddress)
	fmt.Println("Enter the number of tickets: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, emailAddress, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, emailAddress string) {
	remainingTickets -= userTickets

	//create a map for a user

	var userData = UserData{
		firstName:    firstName,
		lastName:     lastName,
		emailAddress: emailAddress,
		numTickets:   userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("%v %v booked %v tickets at email address %v.\n", firstName, lastName, userTickets, emailAddress)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

func sendTickets(userTickets uint, firstName string, lastname string, email string) {
	time.Sleep(10 * time.Second)
	var tickets = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastname)
	fmt.Println("#########")
	fmt.Printf("Sending tickets\n %v \n To email address %v\n", tickets, email)
	fmt.Println("#########")
	wg.Done()
}
