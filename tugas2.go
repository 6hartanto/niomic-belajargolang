package main

import "fmt"

var angka = 7

func main() {
	if angka%2 == 0 {
		fmt.Println("Angka", angka, "adalah angka genap")
	} else {
		fmt.Println("Angka", angka, "adalah angka ganjil")
	}
}
