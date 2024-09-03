package mailService

import (
	"fmt"
	"sync"
	"time"
)

func GenerateTicket(wg *sync.WaitGroup, tickets_needed uint, Fname, email string) {

	time.Sleep(20 * time.Second) //simulating tickect generation and email sending takes some time

	ticket := fmt.Sprintf("User %v has booked %v tickets", Fname, tickets_needed)
	fmt.Println("###################################")
	fmt.Printf(" \n%v\n to %v \n", ticket, email)
	fmt.Println("###################################")
	wg.Done()
}
