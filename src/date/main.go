package main

import (
	"fmt"
	"log"
	"src/calendar"
)

func main() {
	event := calendar.Event{}
	err := event.SetTitle("An extremely long and impractical title")
	if err != nil {
		log.Fatal(err)
	}
	err := event.SetYear(2023) // set-метод типа Date был повышен до Event
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
	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())
}
