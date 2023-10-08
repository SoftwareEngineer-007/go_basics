// tocelsius преобразует температуру из градусов по Фаренгейту в градусы по Цельсию
package main

import (
	"fmt"
	"keyboard"
	"log"
)

func main() {
	fmt.Print("Enter a temperature in Fahrenheit: ")
	fahrenheit, err := keyboard.GetFloat() // вызываем функцию для получения температуры
	if err != nil {
		log.Fatal(err) // если ошибка, выведет сообщение и закроет программу
	}
	celsius := (fahrenheit - 32) * 5 / 9           // преобразование к шкале Цельсия
	fmt.Printf("%0.2f degrees Celsius\n", celsius) // вывод на экран
}
