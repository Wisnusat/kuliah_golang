package main

import "fmt"

func main() {
	var r float64
	var total float64
	fmt.Println("------------------------------")
	fmt.Println("Program Hitung Luas Lingkaran")
	fmt.Println("------------------------------")
	fmt.Print("Masukkan Jari Jari :")
	fmt.Scan(&r)
	total = 3.14 * r * r

	fmt.Println("Luas: ", total)
}
