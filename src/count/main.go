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
	var names []string // хранение сегмента с именами кандидатов
	var counts []int   // сегмент с кол-вом вхождений каждого имени
	for _, line := range lines {
		matched := false
		for i, name := range names { // перебор из сегмента names
			if name == line { // если совпадает с текущим именем...
				counts[i]++    // увеличивает счетчик
				matched = true // устанавливаем признак обнаруженного совпадения
			}
		}
		if matched == false { // если совпадение не найдено...
			names = append(names, line) // добавить его как новое имя в сегмент...
			counts = append(counts, 1)  // и добавить новый счетчик
		}
	}
	for i, name := range names {
		fmt.Printf("%s: %d\n", name, counts[i]) // вывод каждого элемента names и counts
	}
}
