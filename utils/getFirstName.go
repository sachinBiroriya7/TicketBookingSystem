package utils

import "TicketBookingSystem/models"

func GetFirstNames(Booking []models.User) []string {
	f_names := []string{}
	for _, v := range Booking {
		f_names = append(f_names, v.Firstname)
	}
	return f_names
}
