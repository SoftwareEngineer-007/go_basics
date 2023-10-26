package main

import (
	magazine "command-line-argumentsD:\\Projects\\go_basics\\src\\magazine\\magazine.go"
	"fmt"
)

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func printInfo(s *subscriber) { // Объявляется один параметр с типом "subskriber"
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active?", s.active)
}

func defaultSubscriber(name string) *subscriber { // возвращает указатель
	var s subscriber // создается новое значение "subskriber"
	s.name = name
	s.rate = 5.99
	s.active = true
	return &s // возвращает указатель на структуру
}

func applyDiscount(s *subscriber) {
	s.rate = 4.99
}

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
