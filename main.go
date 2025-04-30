package main

import (
	"fmt" 
	"strings"
)

func main() {

	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	// var bookings [50]string
	// var bookingsSlice []string
	// var bookingsSlice = []string{}
	bookingsSlice := []string{}

	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// fmt.Println(conferenceName);
	// fmt.Println(&conferenceName);

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		var fullName string

		// ask user for their name
		fmt.Println("Enter your First Name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your Last Name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your Email:")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets, you want:")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets

		fullName = firstName + " " + lastName
		// bookings[0] = fullName
		// bookingsSlice = append(bookingsSlice, fullName)
		bookingsSlice = append(bookingsSlice, fullName)

		fmt.Println("")
		fmt.Println("**********************")
		fmt.Println("")

		fmt.Printf("Thank you %v %v for booking %v tickets. You will get a confirmation email at %v.\n", firstName, lastName, userTickets, email)

		fmt.Println("")

		fmt.Printf("%v tickets remain out of %v\n", remainingTickets, conferenceTickets)

		fmt.Println("")

		fmt.Printf("No of person make booking: %v \n", len(bookingsSlice))

		fmt.Println("")

		// fmt.Printf("Booking:\n %v \n", bookingsSlice)
		fmt.Println("Booking:")

		fName := ""
		serial := 0

		for index,booking := range bookingsSlice {
			serial = index + 1
			fName = strings.Fields(booking)[0]
			
			fmt.Printf("%v. %v\n", serial, fName)
		}

		fmt.Println("")
		fmt.Println("**********************")
		fmt.Println("")
	}
}
