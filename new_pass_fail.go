// reports whether the user has passed the exam
package main

import (
	"fmt"
	"keyboard"
	"log"
)

func main() {
	fmt.Print("Enter a grade: ")      // запрашиваем у пользователя значение
	grade, err := keyboard.GetFloat() // вызываем функцию из пакета keyboard
	if err != nil {
		log.Fatal(err) // если ошибка, выведет сообщение и закроет программу
	}

	var status string // объявляем переменную здесь, чтоб находилась в области видимости

	if grade >= 60 { // проверка
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status) // вывод введенного значения и результата сдачи
}
