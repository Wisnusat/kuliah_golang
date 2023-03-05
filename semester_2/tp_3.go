package main

import "fmt"

func main() {
	// member()
	// prestasi()
	// gaji()
	tabungan()
}

func member() {
	var total_belanja int
	var status_membership string

	fmt.Scanln(&total_belanja)
	fmt.Scanln(&status_membership)

	if status_membership == "Yes" {
		if total_belanja > 100000 {
			fmt.Println("Anda mendapat discount 5%, tambahan poin 10 dan free gift")
		} else if total_belanja > 75000 {
			fmt.Println("Anda mendapat discount 5%,  tambahan poin 5")
		} else {
			fmt.Println("Anda mendapat tambahan poin 5")
		}
	} else {
		fmt.Println("Total belanja: ", total_belanja)
	}
}

func prestasi() {
	var total1, total2, total3 int
	var predicate string
	total1 = 0
	total2 = 0
	total3 = 0

	fmt.Scanln(&predicate)

	for predicate != "STOP" {
		if predicate == "Sufficient" {
			total1 += 1
		} else if predicate == "Good" {
			total2 += 1
		} else if predicate == "Excellent" {
			total3 += 1
		}
	}
	fmt.Println("Total siswa dengan perdikat Sufficient: ", total1)
	fmt.Println("Total siswa dengan perdikat Good: ", total2)
	fmt.Println("Total siswa dengan perdikat Excellent: ", total3)
}

func gaji() {
	var g string
	var jr, jl, total int

	for g != "Z" {
		fmt.Scanln(&g, &jr, &jl)
		if g == "A" {
			fmt.Println((jr * 75000) + (jl * 90000))
			total += (jr * 75000) + (jl * 90000)
		} else if g == "B" {
			fmt.Println((jr * 125000) + (jl * 70000))
			total += (jr * 125000) + (jl * 70000)
		} else {
			fmt.Println("Hanya mengenal gol A dan B")
		}
	}
	fmt.Println("Biaya yang dikeluarkan perusahaan untuk gaji karyawan", total)
}

func tabungan() {
	var uang, jumlah, tertinggi, terendah, x int

	fmt.Scanln(&uang)
	terendah = uang
	tertinggi = uang
	for uang > 0 {
		if uang > tertinggi {
			tertinggi = uang
		} else if uang < terendah {
			terendah = uang
		}
		jumlah += uang
		x++
		fmt.Scanln(&uang)
	}
	fmt.Println(jumlah)
	fmt.Println(jumlah / x)
	fmt.Println(tertinggi)
	fmt.Println(terendah)
}
