package main

import "fmt"

type Liters float64
type Milliliters float64
type Gallons float64

func LitersToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

func MillilitersToGallons(m Milliliters) Gallons {
	return Gallons(m * 0.000264)
}

func GallonsToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

func GallonsToMilliliters(g Gallons) Milliliters {
	return Milliliters(g * 3785.41)
}

func main() {
	carFuel := Gallons(1.2)
	busFuel := Liters(4.5)
	carFuel += ToGallons(Liters(40.0))
	busFuel += ToLiters(Gallons(30.0))
	fmt.Printf("Car fuel: %0.1f gallons\n", carFuel)
	fmt.Printf("Bus fuel: %0.1f liters\n", busFuel)
}
