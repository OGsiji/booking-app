package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var ConferenceName string = "Go Conference"

const ConferenceTickets int = 50

var remainingTickets uint = 70

var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	FirstName, LastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(FirstName, LastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(userTickets, FirstName, LastName, email)

		wg.Add(1)
		go sendTicket(userTickets, FirstName, LastName, email)

		firstNames := getFirstNames()

		fmt.Printf("first names %v needs to be registered\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")

		}

	} else {
		if !isValidName {
			fmt.Println("first name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("email address you entered doesn't contain @ sign")
		}
		fmt.Printf("The tickets filled are not valid\n")
	}
	wg.Wait()

}

func greetUsers() {
	fmt.Printf("Welcome to %v our Conference\n", ConferenceName)
	fmt.Printf("Welcome to %v booking application\n", ConferenceName)
	fmt.Printf("We have at total of %v tickets and %v are still available.\n", ConferenceName, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")

}

func getFirstNames() []string {

	firstNames := []string{}

	for _, booking := range bookings {

		firstNames = append(firstNames, booking.firstName)

	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {

	var FirstName string
	var LastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first Name: ")
	fmt.Scan(&FirstName)

	fmt.Println("Enter your Last Name: ")
	fmt.Scan(&LastName)

	fmt.Println("Enter your Email Name: ")
	fmt.Scan(&email)

	fmt.Println("Enter your userTickets: ")
	fmt.Scan(&userTickets)

	return FirstName, LastName, email, userTickets
}

func bookTicket(userTickets uint, FirstName string, LastName string, email string) {

	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       FirstName,
		lastName:        LastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// userData["firstName"] = FirstName
	// userData["LastName"] = LastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)

	fmt.Printf("list of bookings is %v bookings\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", FirstName, LastName, userTickets, email)

	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, ConferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###########")
	fmt.Printf("Sending ticket:\n %v \nto email adress %v\n", ticket, email)
	fmt.Println("###########")
	wg.Done()
}
