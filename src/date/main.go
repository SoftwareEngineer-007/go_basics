package main

import "fmt"

type Date struct {
	Year  int
	Month int
	Day   int
}

func main() {
	date := Date{Year: 2023, Month: 10, Day: 29}
	fmt.Println(date)
}
