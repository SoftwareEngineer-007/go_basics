// reports whether the user has passed the exam
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")          // запрашиваем у пользователя значение
	reader := bufio.NewReader(os.Stdin)   // конструкция для чтения ввода с клавиатуры
	input, err := reader.ReadString('\n') // чтение данных
	if err != nil {
		log.Fatal(err) // если ошибка, выведет сообщение и закроет программу
	}

	input = strings.TrimSpace(input)            // удалит символ новой строки
	grade, err := strconv.ParseFloat(input, 64) // преобразует введенную строку в float64
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
