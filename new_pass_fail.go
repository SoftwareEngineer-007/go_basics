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

func getFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}

func main() {
	fmt.Print("Enter a grade: ") // запрашиваем у пользователя значение
	grade, err := getFloat()     // вызываем функцию для получения grade
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
