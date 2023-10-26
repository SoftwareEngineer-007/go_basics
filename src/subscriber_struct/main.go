package main

import (
	magazine "command-line-argumentsD:\\Projects\\go_basics\\src\\magazine\\magazine.go"
	"fmt"
)

func main() {
	subscriber := magazine.Subscriber{Name: "Aman Singh"}
	subscriber.Address.Street = "123 Oak St"
	subscriber.Address.City = "Omaha"
	subscriber.Address.State = "NE"
	subscriber.Address.PostalCode = "68111"
	fmt.Println("Subscriber Name:", subscriber.Name)
	fmt.Println("Street:", subscriber.Address.Street)
	fmt.Println("City:", subscriber.Address.City)
	fmt.Println("State:", subscriber.Address.State)
	fmt.Println("Postal Code:", subscriber.Address.PostalCode)

	employee := magazine.Employee{Name: "Joy Carr"}
	employee.Address.Street = "456 Elm St"
	employee.Address.City = "Portland"
	employee.Address.State = "OR"
	employee.Address.PostalCode = "97222"
	fmt.Println("Employee Name:", employee.Name)
	fmt.Println("Street:", employee.Address.Street)
	fmt.Println("City:", employee.Address.City)
	fmt.Println("State:", employee.Address.State)
	fmt.Println("Postal Code:", employee.Address.PostalCode)
}
