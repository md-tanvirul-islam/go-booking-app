# Go Conference Ticket Booking System

A command-line application for managing ticket bookings for the Go Conference.

## Features

- Interactive ticket booking system
- User input validation (names, email, ticket numbers)
- Concurrent email notifications
- Booking reports and statistics
- Multiple data structure implementations (demonstrating slices, maps, and structs)

## Prerequisites

- Go 1.16 or higher

## Installation

1. Clone the repository:
```bash
git clone [repository-url]
cd booking-app
```

2. Build and run:
```bash
go run main.go
```

## Usage

1. Run the application:
```bash
go run main.go
```

2. Follow the prompts to:
   - Enter your first and last name
   - Provide your email address
   - Specify the number of tickets you want

3. The system will:
   - Validate all inputs
   - Process your booking
   - Send a confirmation email (simulated)
   - Display current booking statistics

## Project Structure

```
booking-app/
├── main.go            # Main application logic
├── helper/            # Helper package for validations
│   └── helper.go      # Validation functions
```

## Key Components

1. **User Struct**:
   ```go
   type User struct {
       FirstName   string
       LastName    string
       Email       string
       UserTickets int
   }
   ```

2. **Validation Functions** (in helper package):
   - `IsValidName()` - Validates first and last names
   - `IsValidEmail()` - Validates email format
   - `IsValidTicketNumber()` - Validates ticket quantity

3. **Concurrent Operations**:
   - Uses goroutines and WaitGroup for simulated email sending

## Data Structure Options

The code demonstrates three approaches (commented out):
1. **Slices** - Simple string storage
2. **Maps** - Key-value pair storage
3. **Structs** (active implementation) - Type-safe structured data

## Example Session

```
Welcome to Go Conference booking application
We have total of 50 tickets and 50 are still available.
Get your tickets here to attend

Enter your First Name:
John
Enter your Last Name:
Doe
Enter your Email:
john@example.com
Enter number of tickets you want:
2

******* Report **********

Thank you John Doe for booking 2 tickets. You will get a confirmation email at john@example.com.

48 tickets remain out of 50

No of person make booking: 1 

Booking:
1. John

*************************
```

## Notes

- The email sending is simulated with a 30-second delay
- Application exits when all tickets are booked
- Clean, modular code with separation of concerns