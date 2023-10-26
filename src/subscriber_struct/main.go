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
	address := magazine.Address{Street: "123 Oak St", City: "Omaha", State: "NE", PostalCode: "68111"}
	subscriber := magazine.Subscriber{Name: "Aman Singh"}
	subscriber.HomeAddress = address
	fmt.Println(subscriber.HomeAddress)
}
