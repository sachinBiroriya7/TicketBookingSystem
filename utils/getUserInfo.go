package utils

import "fmt"

func GetUserInfo() (uint, string, string, string) {
	//we dont need any args from main to this func, but need to return value to main to do further processing

	var FName string
	var LName string
	var tickets_needed uint
	var email string

	fmt.Println("enter your first name : ")
	fmt.Scan(&FName)

	fmt.Println("enter your last name : ")
	fmt.Scan(&LName)

	fmt.Println("enter number of tickets required :")
	fmt.Scan(&tickets_needed)

	fmt.Println("enter you email id :")
	fmt.Scan(&email)

	return tickets_needed, FName, LName, email
}
