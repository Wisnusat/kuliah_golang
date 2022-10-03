package main

import "fmt"

func soal1() {
	var x, y float64
	fmt.Println("=====================")
	fmt.Println("Soal 1")
	fmt.Println("=====================")
	fmt.Print("Masukkan nilai X: ")
	fmt.Scanln(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scanln(&y)
	var hasil float64 = 1/(3*x*x) + (10 * y) + 7
	fmt.Println("Hasil :", hasil)
	fmt.Println("---------------------")
	fmt.Println("")
}

func soal2() {
	var x float64
	fmt.Println("=====================")
	fmt.Println("Soal 2")
	fmt.Println("=====================")
	fmt.Print("Masukkan nilai X: ")
	fmt.Scanln(&x)
	var hasil float64 = ((x * x * x) + (3 * x)) / ((x * x * x * x) - (3 * x * x) + 4)
	fmt.Println("Hasil :", hasil)
	fmt.Println("---------------------")
	fmt.Println("")
}

func soal3() {
	var x, d1, d2, d3 int
	fmt.Println("=====================")
	fmt.Println("Soal 3")
	fmt.Println("=====================")
	fmt.Print("Masukkan nilai X: ")
	fmt.Scanln(&x)
	d3 = x % 10
	d2 = ((x - d3) / 10) % 10
	d1 = ((x - (d3 + (d2 * 10))) / 100) % 10
	fmt.Println("d1 :", d1)
	fmt.Println("d2 :", d2)
	fmt.Println("d3 :", d3)
	fmt.Println("---------------------")
	fmt.Println("")

}

func main() {
	soal1()
	soal2()
	soal3()
}
