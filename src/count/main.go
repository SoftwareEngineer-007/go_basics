// count подсчитывает количество вхождений каждой строки в файле
package main

import (
	"datafile"
	"fmt"
	"log"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int) // объявляем карту
	for _, line := range lines {
		counts[line]++ // увеличиваем счетчик голосов
	}
	fmt.Println(counts) // выводит заполненную карту
}
