package main

import (
	"fmt"
	"strconv"
	"time"

	// "strings"
	"booking-app/helper"
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
	// bookingsSlice := []string{}
	// bookingsMap := make([]map[string]string, 0)
	bookingsStruct := make([]User, 0)

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
		var firstName string
		var lastName string
		var email string
		var userTicketsInput string
		// var fullName string
		var userTickets int

		fmt.Println("Enter your First Name:")
		fmt.Scan(&firstName)

		if !helper.IsValidName(firstName, "First name") {
			continue
		}

		fmt.Println("Enter your Last Name:")
		fmt.Scan(&lastName)

		if !helper.IsValidName(lastName, "Last name") {
			continue
		}

		fmt.Println("Enter your Email:")
		fmt.Scan(&email)

		if !helper.IsValidEmail(email) {
			continue
		}

		fmt.Println("Enter number of tickets, you want:")
		fmt.Scan(&userTicketsInput)

		if !helper.IsValidTicketNumber(userTicketsInput) {
			continue
		}

		userTickets, _ = strconv.Atoi(userTicketsInput)

		if userTickets > remainingTickets {
			overTicketBookingWarning(remainingTickets, userTickets)

		} else {
			remainingTickets = remainingTickets - userTickets

			/* Using Slice */
			// fullName = firstName + " " + lastName
			// bookingsSlice = append(bookingsSlice, fullName)
			// bookingReport(user, remainingTickets, conferenceTickets, bookingsSlice)

			/* Using Map */
			// var userMap = make(map[string]string)
			// userMap["firstName"] = firstName
			// userMap["lastName"] = lastName
			// userMap["email"] = email
			// userMap["userTickets"] = strconv.FormatInt(int64(userTickets), 10)
			// bookingsMap = append(bookingsMap, userMap)
			// bookingReport(user, remainingTickets, conferenceTickets, bookingsMap)

			/* Using Struct */
			user := User{
				FirstName:   firstName,
				LastName:    lastName,
				Email:       email,
				UserTickets: userTickets,
			}

			bookingsStruct = append(bookingsStruct, user)
			bookingReport(user, remainingTickets, conferenceTickets, bookingsStruct)
			go emailBooking(user.FirstName, user.Email)

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

// func bookingReport(user User, remainingTickets int, conferenceTickets int, bookings []string) { // using slices
// func bookingReport(user User, remainingTickets int, conferenceTickets int, bookings []map[string]string) { // using list of map
func bookingReport(user User, remainingTickets int, conferenceTickets int, bookings []User) { // using list of struct
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

	// fName := ""
	serial := 0

	for index, booking := range bookings {
		serial = index + 1

		/* Using Slice */
		// fName = strings.Fields(booking)[0]
		// fmt.Printf("%v. %v\n", serial, fName)

		/* Using Map */
		// fmt.Printf("%v. %v\n", serial, booking["firstName"])

		/* Using Struct */
		fmt.Printf("%v. %v\n", serial, booking.FirstName)
	}

	fmt.Println("")
	fmt.Println("*************************")
	fmt.Println("")
}

func overTicketBookingWarning(remainingTickets int, userTickets int) {
	fmt.Println("")
	fmt.Println("******* Warning **************")
	fmt.Printf("We have %v tickets remaining. You can't book %v tickets\n", remainingTickets, userTickets)
	fmt.Println("******************************")
}

func emailBooking(firstName string, email string) {
	time.Sleep(30 * time.Second)

	body := fmt.Sprintf("Hi, %v You have successfully booked the ticket for the Conference.\n", firstName)
	fmt.Println("******* Emailing **************")
	fmt.Printf("Address: %v \nBody: %v\n", email, body)
	fmt.Println("*******************************")
}
