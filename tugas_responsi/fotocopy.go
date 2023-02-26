package main

import "fmt"

func main() {
	var isStudent, isSkripsi bool
	var berlangganan, priority int
	var total float64
	fmt.Scanln(&isStudent, &isSkripsi, &berlangganan, &total)
	if isStudent && berlangganan > 12 {
		priority = 1
		if isSkripsi {
			fmt.Println("Pelanggan prioritas", priority)
			fmt.Println(total - (total * 35 / 100))
		} else {
			fmt.Println("Pelanggan prioritas", priority)
			fmt.Println(total - (total * 30 / 100))
		}
	} else if isStudent && berlangganan > 8 && berlangganan <= 12 {
		priority = 2
		if isSkripsi {
			fmt.Println("Pelanggan prioritas", priority)
			fmt.Println(total - (total * 20 / 100))
		} else {
			fmt.Println("Pelanggan prioritas", priority)
			fmt.Println(total - (total * 15 / 100))
		}
	} else if isStudent {
		fmt.Println(total - (total * 10 / 100))
	} else {
		fmt.Println(total)
	}
}
