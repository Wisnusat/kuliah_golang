package main

import "fmt"

func main() {
	var jenis string
	fmt.Scanln(&jenis)
	if jenis == "mudah" {
		fmt.Println("Keren! Good job! :D")
	} else if jenis == "sedang" {
		fmt.Println("Nice try :D")
	} else if jenis == "sulit" {
		fmt.Println("yu bisa yu")
	} else {
		fmt.Println("Inputan tidak valid")
	}
}
