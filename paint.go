package main

import "fmt"

func paitNeeded(width float64, height float64) {
	area := width * height
	fmt.Printf("%.2f liters needed\n", area/10.0)
}

func main() {
	paitNeeded(4.2, 3.0)
	paitNeeded(5.2, 3.5)
	paitNeeded(5.0, 3.3)
}
