package main

import (
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	FirstName   string
	LastName    string
	Email       string
	UserTickets int
}

func main() {

	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50
	bookingsSlice := []string{}

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
		var firstName string
		var lastName string
		var email string
		var userTicketsInput string
		var fullName string
		var userTickets int

		fmt.Println("Enter your First Name:")
		fmt.Scan(&firstName)

		if !isValidName(firstName, "First name") {
			continue
		}

		fmt.Println("Enter your Last Name:")
		fmt.Scan(&lastName)

		if !isValidName(lastName, "Last name") {
			continue
		}

		fmt.Println("Enter your Email:")
		fmt.Scan(&email)

		if !isValidEmail(email) {
			continue
		}

		fmt.Println("Enter number of tickets, you want:")
		fmt.Scan(&userTicketsInput)

		if !isValidTicketNumber(userTicketsInput) {
			continue
		}

		userTickets, _ = strconv.Atoi(userTicketsInput)

		if userTickets > remainingTickets {
			overTicketBookingWarning(remainingTickets, userTickets)

		} else {
			remainingTickets = remainingTickets - userTickets

			fullName = firstName + " " + lastName
			bookingsSlice = append(bookingsSlice, fullName)

			user := User{
				FirstName:   firstName,
				LastName:    lastName,
				Email:       email,
				UserTickets: userTickets,
			}

			bookingReport(user, remainingTickets, conferenceTickets, bookingsSlice)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked. Come back next year.")
				break
			}
		}

	}
}

func greetUsers(conferenceName string, conferenceTickets int, remainingTickets int) {
	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func bookingReport(user User, remainingTickets int, conferenceTickets int, bookings []string) {
	fmt.Println("")
	fmt.Println("******* Report **********")
	fmt.Println("")

	fmt.Printf("Thank you %v %v for booking %v tickets. You will get a confirmation email at %v.\n", user.FirstName, user.LastName, user.UserTickets, user.Email)

	fmt.Println("")

	fmt.Printf("%v tickets remain out of %v\n", remainingTickets, conferenceTickets)

	fmt.Println("")

	fmt.Printf("No of person make booking: %v \n", len(bookings))

	fmt.Println("")

	// fmt.Printf("Booking:\n %v \n", bookingsSlice)
	fmt.Println("Booking:")

	fName := ""
	serial := 0

	for index, booking := range bookings {
		serial = index + 1
		fName = strings.Fields(booking)[0]

		fmt.Printf("%v. %v\n", serial, fName)
	}

	fmt.Println("")
	fmt.Println("******* Report **********")
	fmt.Println("")
}

func isValidEmail(email string) bool {
	var valid bool = true

	if !(strings.Contains(email, "@") && strings.Contains(email, ".")) {
		fmt.Println("")
		fmt.Println("******* Warning **************")
		fmt.Println("Email is not valid.")
		fmt.Println("")
		valid = false
	}

	return valid
}

func isValidName(lastName string, fieldName string) bool {
	var valid bool = true

	if len(lastName) <= 2 {
		fmt.Println("")
		fmt.Println("******* Warning **************")
		fmt.Printf("%v is too short.\n", fieldName)
		fmt.Println("")
		valid = false
	}
	return valid
}

func isValidTicketNumber(userTicketsInput string) bool {
	var valid bool = true

	userTickets, userTicketsInputErr := strconv.Atoi(userTicketsInput)

	if userTicketsInputErr != nil && userTickets < 1 {
		fmt.Println("")
		fmt.Println("******* Warning **************")
		fmt.Println("Invalid ticket number.")
		fmt.Println("")
		valid = false
	}
	return valid
}

func overTicketBookingWarning(remainingTickets int, userTickets int) {
	fmt.Println("")
	fmt.Println("******* Warning **************")
	fmt.Printf("We have %v tickets remaining. You can't book %v tickets\n", remainingTickets, userTickets)
	fmt.Println("")
}
