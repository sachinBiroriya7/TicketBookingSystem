package main

import (
	"TicketBookingSystem/mailService"
	"TicketBookingSystem/models"
	"TicketBookingSystem/utils"
	"fmt"
	"strings"
	"sync"
)

const Total_tickets uint = 50

var remaining_tickets uint = 50

var Booking []models.User

var wg = sync.WaitGroup{}

func welcomeUser(Total_tickets, remaining_tickets uint) {
	fmt.Println("*******************Hi, Welcome to the Conference Ticket Booking system*******************")
	fmt.Printf("We had %v tickets, and We are left with %v tickets, Continue Booking details...\n", Total_tickets, remaining_tickets)
}

func bookTickets(tickets_needed uint, FName, LName, email string) {
	fmt.Printf("Congrats %v %v, You've booked %v tickets, confirmation will be sent to %v\n", FName, LName, tickets_needed, email)
	remaining_tickets = remaining_tickets - tickets_needed

	fmt.Printf("Remaining tickets are : %v \n", remaining_tickets)

	var UserData = models.User{
		Firstname:     FName,
		LastName:      LName,
		Email:         email,
		BookedTickets: tickets_needed,
	}

	Booking = append(Booking, UserData)
	fmt.Printf("All booking in our Conference are %v\n", Booking)
}

func main() {

	welcomeUser(Total_tickets, remaining_tickets)

	for remaining_tickets > 0 {

		tickets_needed, FName, LName, email := utils.GetUserInfo()

		isValidEmail, isValidFName, isValidLName, isvalidTickets := validateUserInput(tickets_needed, Total_tickets, remaining_tickets, FName, LName, email)

		if isvalidTickets && isValidFName && isValidLName && isValidEmail {

			bookTickets(tickets_needed, FName, LName, email)

			wg.Add(1)                                                        //added a goroutine
			go mailService.GenerateTicket(&wg, tickets_needed, FName, email) //pasing &wg to mailService package

			fname := utils.GetFirstNames(Booking)
			fmt.Println("First Name := ", fname)

			var flag string
			fmt.Print("Do you want to book again [Y/N]? ")
			fmt.Scan(&flag)

			flag = strings.ToLower(flag)
			if flag != "y" {
				break
			}
		} else {
			if !isValidFName || !isValidLName {
				fmt.Println("entered first or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("check email, it doesn't contain `@` or `.`")
			}
			if !isvalidTickets {
				fmt.Printf("Cant book %v tickets, as we are left with %v tickets only\n", tickets_needed, remaining_tickets)
			}

		}

	}
	wg.Wait()
}
