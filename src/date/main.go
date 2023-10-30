package main

import (
	"fmt"
	"log"
	"src/calendar"
)

func main() {
	event := calendar.Event{}
	err := event.SetTitle("Mom's birthday") // определяется для типа Event
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetYear(2023) // set-метод типа Date был повышен до Event
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(10) // set-метод типа Date был повышен до Event
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(29) // set-метод типа Date был повышен до Event
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Title()) // определяется для типа Event
	fmt.Println(event.Year())  // определяется для типа Date
	fmt.Println(event.Month()) // определяется для типа Date
	fmt.Println(event.Day())   // определяется для типа Date
}
