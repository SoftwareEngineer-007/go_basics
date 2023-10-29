package main

import "fmt"

type Liters float64
type Milliliters float64
type Gallons float64

func (l Liters) ToGallons() Gallons { // определяем метод ToGallons для типа Liters
	return Gallons(l * 0.264)
}

func (m Milliliters) ToGallons() Gallons { // определяем метод ToGallons для типа Milliliters
	return Gallons(m * 0.000264)
}

func (g Gallons) ToLiters() Liters { // определяем метод ToLiters для типа Gallons
	return Liters(g * 3.785)
}

func (g Gallons) ToMilliliters() Milliliters { // определяем метод ToMilliliters для типа Gallons
	return Milliliters(g * 3785.41)
}

func main() {
	soda := Liters(2)                                                                // создание значения Liters
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", soda, soda.ToGallons())        // Преобразование Liters в Gallons
	water := Milliliters(500)                                                        // создание значения Milliliters
	fmt.Printf("%0.3f milliliters equals %0.3f gallons\n", water, water.ToGallons()) // Преобразование Milliliters в Gallons

	milk := Gallons(2)                                                                 // Создаем значение Gallons
	fmt.Printf("%0.3f gallons equals %0.3f liters\n", milk, milk.ToLiters())           // Преобразуем в Liters
	fmt.Printf("%0.3f gallons equals %0.3f milliliters\n", milk, milk.ToMilliliters()) // Преобразуем в Milliliters
}
