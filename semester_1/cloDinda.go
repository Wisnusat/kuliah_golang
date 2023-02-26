package main

import (
	"fmt"
	"math"
)

func soal1() {
	var bil, jumlah int
	jumlah = 1
	fmt.Scanln(&bil)
	for bil > 0 {
		jumlah = jumlah * (bil % 10)
		bil = bil / 10
	}
	fmt.Println(jumlah)
}

func soal2() {
	var bil, jumlah int
	fmt.Scanln(&bil)
	for bil > 0 {
		jumlah = jumlah + (bil % 10)
		bil = bil / 10
	}
	fmt.Println(jumlah)
}

func soal3() {
	var bil int
	fmt.Scanln(&bil)
	for bil > 0 {
		fmt.Print(bil % 10)
		bil = bil / 10
	}
}

func soal7() {
	var bacteriaPresent, number, time, hari int
	fmt.Scanln(&bacteriaPresent, &time)
	for hari = 1; hari <= time; hari++ {
		number = bacteriaPresent * int(math.Pow(float64(2), float64(hari/10)))
		fmt.Println(number, "selama", hari, "hari")
	}
}

func coba() {
	var jumlah float64
	jumlah = math.Pow(2, 3)
	fmt.Println(jumlah)
}

func soal5() {
	var n, i, hasil int
	fmt.Scan(&n)
	hasil = 1
	if n > 0 {
		for i = 1; i <= n; i++ {
			hasil *= i
		}
		fmt.Println(hasil)
	} else {
		fmt.Println("Inputan harus bilangan bulat positif")
	}
}

func soal4() {
	var n, i, jumlah, bil, rata_rata float64
	fmt.Print("Jumlah item: ")
	fmt.Scanln(&n)
	if n > 0 {
		for i = 1; i <= n; i++ {
			fmt.Print("Bilangan ke-", i, ": ")
			fmt.Scanln(&bil)
			jumlah += bil
		}
		rata_rata = jumlah / n
		fmt.Println("Rata-rata: ", rata_rata)
	} else {
		fmt.Println("Jumlah item harus lebih dari 0")
	}

}

func main() {
	// soal1()
	// soal2()
	// soal3()
	// soal7()
	// coba()
	// soal5()
	soal4()
}
