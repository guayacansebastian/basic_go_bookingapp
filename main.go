package main

import (
	"fmt"
	"strings"
)

const tickets = 50          //const, int
var conference = "MY EVENT" //var, string
var remaining int = 50.0    //var, int
var userName string
var lastName string
var email string
var userTickets int

func main() {
	// Declare variables and constants

	// var userName string         // var, string
	// var userTickets int
	// userTickets := 0 // shorthand (var), float64

	// Prints
	// fmt.Println("Welcome to", conference, "booking application")
	// fmt.Printf("Get your tickets for the  %v event. \nWe have %v tickets  and %v are available.\n", conference, tickets, remaining)

	greetings()
	// fmt.Printf("UserName is %T, usertickts is %T\n", userName, userTickets) // %T prints the type of variable
	// fmt.Printf("Pointer example: %v\n", &userTickets)                       // & gives the memory address of variable
	// user input
	// fmt.Println("Enter your name:")
	// fmt.Scan(&userName) // & gives the memory address of variable

	// var lastName string
	// fmt.Println("Enter your last name:")
	// fmt.Scan(&lastName)

	// var email string
	// fmt.Println("Enter your email:")
	// fmt.Scan(&email)

	// fmt.Println("Enter number of tickets:")
	// fmt.Scan(&userTickets)
	// remaining = remaining - userTickets

	// // arrays
	// var bookings [50]string // array of strings
	// bookings[0] = userName + " " + lastName
	// fmt.Printf("Complete array of bookings: %v\n", bookings)
	// fmt.Printf("First booking: %v\n", bookings[0])
	// fmt.Printf("Number of bookings: %v\n", len(bookings))
	// slices
	var bookingsSlice []string // slice of strings
	// bookingsSlice = append(bookingsSlice, userName+" "+lastName)
	// fmt.Printf("Slice of bookings: %v\n", bookingsSlice)
	// fmt.Printf("Number of bookings in slice: %v\n", len(bookingsSlice))

	//loops
	for {
		userName, lastName, email, userTickets := get_user_input()

		validname, validmail, validnumber := validadateuserinput(userName, lastName, email, userTickets, remaining)
		//if condition
		if validmail && validname && validnumber {
			remaining, bookingsSlice = book_ticket(remaining, userTickets, userName, lastName, email, conference, bookingsSlice)

		} else {
			fmt.Println("Invalid input. Please try again.")

			if !validname {
				fmt.Println("Invalid name. Please enter a valid full name with at least 2 characters.")
			}
			if !validmail {
				fmt.Println("Invalid email. Please enter a valid email address.")
			}
			if !validnumber {
				fmt.Printf("Invalid number of tickets. Please enter a number between 1 and %v.\n", remaining)
			}

			continue // skip the rest of the loop and start over
			// break
		}

		// if condition
		fmt.Printf("Remaining tickets: %v\n", remaining)
		if remaining <= 0 {
			fmt.Println("All tickets are sold out!")
			break // exit the loop
		}
	}
	//switch statement
	// city := "New York"
	// switch city {
	// case "New York":
	// 	fmt.Println("You are in New York City!")
	// case "Los Angeles":
	// 	fmt.Println("You are in Los Angeles!")
	// case "Chicago":
	// 	fmt.Println("You are in Chicago!")
	// default:
	// 	fmt.Println("City not recognized.")
	// }
}
func greetings() {

	fmt.Println("Hello, wellcome to the application")
	fmt.Printf("You can buy your tickets for %v event. There are %v tickets available.\n", conference, tickets) //We are using global variable here
}
func firstnames(bookings []string) []string {
	firstnames := []string{}
	for _, booking := range bookings { // range iterates over the array, range returns 2 values: index and value, if we don't need the index, we can use the underscore _
		names := strings.Fields(booking)
		firstname := names[0]
		firstnames = append(firstnames, firstname)
	}
	return firstnames
}

func validadateuserinput(username string, lastname string, email string, usertickets int, remainingtickets int) (bool, bool, bool) {
	validname := len(username) >= 2 && len(lastname) >= 2
	validmail := strings.Contains(email, "@") && strings.Contains(email, ".com")
	validnumber := usertickets <= remainingtickets && usertickets > 0
	return validname, validmail, validnumber

}

func get_user_input() (string, string, string, int) {

	fmt.Println("Enter your name:")
	fmt.Scan(&userName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email:")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)
	return userName, lastName, email, userTickets

}
func book_ticket(remaining int, userTickets int, userName string, lastName string, email string, conference string, bookingsSlice []string) (int, []string) {
	remaining = remaining - userTickets
	fmt.Printf("Thank you %v for booking %v tickets.\n", userName, userTickets)
	fmt.Printf("Slice of bookings: %v\n", bookingsSlice)
	bookingsSlice = append(bookingsSlice, userName+" "+lastName)
	firstnames := firstnames(bookingsSlice)
	fmt.Printf("First names of all bookings: %v\n", firstnames)

	fmt.Printf("Thank you %v for booking %v tickets.\n", userName, userTickets)
	fmt.Printf("%v tickets remaining for %v\n", remaining, conference)
	fmt.Printf("You will receive a confirmation email at %v\n", email)
	return remaining, bookingsSlice

}
