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
	var address magazine.Address
	address.Street = "123 Oak St"
	address.City = "Omaha"
	address.State = "NE"
	address.PostalCode = "68111"
	fmt.Println(address)
}
