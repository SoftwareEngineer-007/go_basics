package main

import (
	"fmt"
	"log"
	"src/calendar"
)

func main() {
	date := calendar.Date{}
	err := date.SetYear(2023) // испоьзуется set-метод
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(10) // испоьзуется set-метод
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(29) // испоьзуется set-метод
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)
}
