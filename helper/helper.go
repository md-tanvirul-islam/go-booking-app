package helper

import (
	"fmt"
	"strings"
	"strconv"
)

func IsValidEmail(email string) bool {
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

func IsValidName(lastName string, fieldName string) bool {
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

func IsValidTicketNumber(userTicketsInput string) bool {
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