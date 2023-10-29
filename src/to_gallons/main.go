package main

import "fmt"

type Liters float64
type Milliliters float64
type Gallons float64

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func main() {
	soda := Liters(2)                                                                // создание значения Liters
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", soda, soda.ToGallons())        // Преобразование Liters в Gallons
	water := Milliliters(500)                                                        // создание значения Milliliters
	fmt.Printf("%0.3f milliliters equals %0.3f gallons\n", water, water.ToGallons()) // Преобразование Milliliters в Gallons
}
