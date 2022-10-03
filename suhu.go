package main

import (
	"fmt"
	"strings"
)

func main() {
	var suhu float64
	var jenis string
	var convert string
	fmt.Print("Masukkan suhu :")
	fmt.Scanln(&suhu)
	fmt.Print("Masukkan jenis suhu :")
	fmt.Scanln(&jenis)
	fmt.Print("Konversi ke :")
	fmt.Scanln(&convert)

	var jenisSuhu string = strings.ToLower(jenis)
	var convertTo string = strings.ToLower(convert)

	// rumus
	if jenisSuhu == "c" || jenisSuhu == "celcius" {
		if convertTo == "r" || convertTo == "reamur" {
			var hasil float64 = suhu * 4 / 5
			fmt.Println(hasil, "R")
		} else if convertTo == "f" || convertTo == "farenheit" {
			var hasil float64 = suhu*9/5 + 32
			fmt.Println(hasil, "F")
		} else {
			fmt.Println("Jenis suhu tidak valid")
		}
	} else if jenisSuhu == "r" || jenisSuhu == "reamur" {
		if convertTo == "c" || convertTo == "celcius" {
			var hasil float64 = suhu * 5 / 4
			fmt.Println(hasil, "C")
		} else if convertTo == "f" || convertTo == "farenheit" {
			var hasil float64 = suhu*9/4 + 32
			fmt.Println(hasil, "F")
		} else {
			fmt.Println("Jenis suhu tidak valid")
		}
	} else if jenisSuhu == "f" || jenisSuhu == "farenheit" {
		if convertTo == "c" || convertTo == "celcius" {
			var hasil float64 = (suhu - 32) * 5 / 9
			fmt.Println(hasil, "C")
		} else if convertTo == "r" || convertTo == "reamur" {
			var hasil float64 = (suhu - 32) * 4 / 9
			fmt.Println(hasil, "R")
		} else {
			fmt.Println("Jenis suhu tidak valid")
		}
	} else {
		fmt.Println("Jenis suhu tidak valid")
	}
}
