// maximum выводит наибольшее число
package main

import (
	"fmt"
	"math"
)

func maximum(numbers ...float64) float64 { // получает любое количество аргументов
	max := math.Inf(-1)              // начинает с очень низкого значения
	for _, number := range numbers { // обработка всех аргументов
		if number > max {
			max = number // находит наибольшее значение
		}
	}
	return max
}

func main() {
	fmt.Println(maximum(71.8, 56.2, 89.5))
	fmt.Println(maximum(90.7, 89.7, 98.5, 92.3))
}
