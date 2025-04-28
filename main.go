package main

import "fmt"

func main() {

	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// fmt.Println(conferenceName);
	// fmt.Println(&conferenceName);

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	var bookings [50]string

	// ask user for their name
	fmt.Println("Enter your First Name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your Last Name:")
	fmt.Scan(&lastName)
	
	fmt.Println("Enter your Email:")
	fmt.Scan(&email)
	
	fmt.Println("Enter number of tickets, you want:")
	fmt.Scan(&userTickets)
	
	remainingTickets = remainingTickets - userTickets;

	bookings[0] = firstName + " " + lastName
	
	fmt.Println("")
	fmt.Println("**********************")
	fmt.Println("")

	fmt.Printf("Thank you %v %v for booking %v tickets. You will get a confirmation email at %v.\n", firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets remain out of %v\n", remainingTickets, conferenceTickets);

	fmt.Printf("Booking:\n %v \n", bookings);
	fmt.Printf("No of person make booking: %v \n", len(bookings));

	fmt.Println(bookings);
	// fmt.Println(bookings[2]);
	// fmt.Println(bookings[3]);

	fmt.Println("")
	fmt.Println("**********************")
	fmt.Println("")
}
