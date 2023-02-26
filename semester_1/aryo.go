package main

import "fmt"

func kue() {
	var kue, orang, jumlah int

	fmt.Scanln(&kue, &orang)
	for kue >= orang {
		jumlah += 1
		kue -= orang
	}
	fmt.Println("Tiap orang dapat", jumlah)
	fmt.Println("Sisa kue", kue)
}

// func rata_rata() {
// 	var n, i, jumlah, bil, rata_rata float64
// 	fmt.Scanln(&n)
// 	if n > 0 {
// 		for i = 1; i <= n; i++ {
// 			fmt.Scanln(&bil)
// 			jumlah += bil
// 		}
// 		rata_rata = jumlah / n
// 		fmt.Println(rata_rata)
// 	} else {
// 		fmt.Println("Jumlah item harus lebih dari 0")
// 	}
// }

func rata_rata() {
	var silver, gold, platinum, gold_user, silver_user, platinum_user, bil, total float64
	for total < 500.0 {
		fmt.Scanln(&bil)
		if bil >= 50.0 && bil <= 99.0 {
			silver += bil
			silver_user += 1.0
			total += bil
		} else if bil >= 100.0 && bil <= 200.0 {
			platinum += bil
			platinum_user += 1.0
			total += bil
		} else if bil > 200.0 {
			gold += bil
			gold_user += 1.0
			total += bil
		}
	}
	fmt.Println("Rata-rata gold user:", gold/gold_user)
	fmt.Println("Rata-rata platinum user:", platinum/platinum_user)
	fmt.Println("Rata-rata silver user:", silver/silver_user)
}

func main() {
	// kue()
	rata_rata()
}
